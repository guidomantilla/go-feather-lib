package messaging

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type RabbitMQConsumerOption func(*RabbitMQConsumer)

func WithAutoAck(autoAck bool) RabbitMQConsumerOption {
	return func(queue *RabbitMQConsumer) {
		queue.autoAck = autoAck
	}
}

func WithNoLocal(noLocal bool) RabbitMQConsumerOption {
	return func(queue *RabbitMQConsumer) {
		queue.noLocal = noLocal
	}
}

func WithDurable(durable bool) RabbitMQConsumerOption {
	return func(queue *RabbitMQConsumer) {
		queue.durable = durable
	}
}

func WithAutoDelete(autoDelete bool) RabbitMQConsumerOption {
	return func(queue *RabbitMQConsumer) {
		queue.autoDelete = autoDelete
	}
}

func WithExclusive(exclusive bool) RabbitMQConsumerOption {
	return func(queue *RabbitMQConsumer) {
		queue.exclusive = exclusive
	}
}

func WithNoWait(noWait bool) RabbitMQConsumerOption {
	return func(queue *RabbitMQConsumer) {
		queue.noWait = noWait
	}
}

func WithArgs(args amqp.Table) RabbitMQConsumerOption {
	return func(queue *RabbitMQConsumer) {
		queue.args = args
	}
}

func WithRabbitMQListener(listener Listener[*amqp.Delivery]) RabbitMQConsumerOption {
	return func(consumer *RabbitMQConsumer) {
		consumer.listener = listener
	}
}

type RabbitMQConsumer struct {
	connection Connection[*amqp.Connection]
	listener   Listener[*amqp.Delivery]
	channel    *amqp.Channel
	queue      amqp.Queue
	name       string
	consumer   string
	autoAck    bool
	noLocal    bool
	durable    bool
	autoDelete bool
	exclusive  bool
	noWait     bool
	args       amqp.Table
	mu         sync.RWMutex
}

func NewRabbitMQConsumer(connection Connection[*amqp.Connection], name string, options ...RabbitMQConsumerOption) *RabbitMQConsumer {

	if connection == nil {
		log.Fatal("starting up - error setting up rabbitmq consumer: connection is nil")
	}

	if strings.TrimSpace(name) == "" {
		log.Fatal("starting up - error setting up rabbitmq consumer: name is empty")
	}

	consumer := &RabbitMQConsumer{
		connection: connection,
		listener:   NewRabbitMQListener(),
		name:       name,
		consumer:   "consumer-" + name,
		autoAck:    false,
		noLocal:    false,
		durable:    false,
		autoDelete: false,
		exclusive:  false,
		noWait:     false,
		args:       nil,
	}

	for _, option := range options {
		option(consumer)
	}
	consumer.autoAck = false

	return consumer
}

func (consumer *RabbitMQConsumer) Consume(ctx context.Context) (Event, error) {

	consumer.mu.Lock()
	defer consumer.mu.Unlock()

	var err error
	var connection *amqp.Connection
	if connection, err = consumer.connection.Connect(); err != nil {
		log.Debug(fmt.Sprintf("rabbitmq consumer - failed connection to queue %s: %s", consumer.name, err.Error()))
		return nil, err
	}

	if !(consumer.channel != nil && !consumer.channel.IsClosed()) {
		if consumer.channel, err = connection.Channel(); err != nil {
			log.Debug(fmt.Sprintf("rabbitmq consumer - failed connection to queue %s: %s", consumer.name, err.Error()))
			return nil, err
		}
	}

	if consumer.queue, err = consumer.channel.QueueDeclare(consumer.name, consumer.durable, consumer.autoDelete, consumer.exclusive, consumer.noWait, consumer.args); err != nil {
		log.Debug(fmt.Sprintf("rabbitmq consumer - failed connection to queue %s: %s", consumer.name, err.Error()))
		return nil, err
	}

	log.Debug(fmt.Sprintf("rabbitmq consumer - connected to queue %s", consumer.name))

	var deliveries <-chan amqp.Delivery
	if deliveries, err = consumer.channel.ConsumeWithContext(ctx, consumer.name, consumer.consumer, consumer.autoAck, consumer.exclusive, consumer.noLocal, consumer.noWait, consumer.args); err != nil {
		log.Debug(fmt.Sprintf("rabbitmq consumer - failed comsuming from queue: %s", err.Error()))
		return nil, err
	}

	closeChannel := make(chan string)
	closeHandler := func(ctx context.Context, listener Listener[*amqp.Delivery], channel *amqp.Channel, queue string, closeChannel chan string) {
		var err error
		for message := range deliveries {
			go func(ctx context.Context, message amqp.Delivery) {
				log.Debug(fmt.Sprintf("rabbitmq consumer - message received: %s", message.Body))
				if err := listener.OnMessage(ctx, &message); err != nil {
					log.Debug(fmt.Sprintf("rabbitmq consumer - failed to process message: %s", err.Error()))
					if err := message.Nack(false, true); err != nil {
						log.Debug(fmt.Sprintf("rabbitmq consumer - failed to nack message: %s", err.Error()))
					}
					log.Debug(fmt.Sprintf("rabbitmq consumer - nack message: %s", err.Error()))
					return
				}
				if err := message.Ack(false); err != nil {
					log.Debug(fmt.Sprintf("rabbitmq consumer - failed to ack message: %s", err.Error()))
					return
				}
				log.Debug(fmt.Sprintf("rabbitmq consumer - ack message: %s", err.Error()))
			}(ctx, message)
		}
		if err = channel.Close(); err != nil {
			log.Debug(fmt.Sprintf("rabbitmq consumer - failed to close channel to queue %s: %s", queue, err.Error()))
			return
		}
		close(closeChannel)
		log.Debug(fmt.Sprintf("rabbitmq consumer - disconected from queue %s", queue))
	}

	go closeHandler(ctx, consumer.listener, consumer.channel, consumer.name, closeChannel)
	return closeChannel, nil
}

func (consumer *RabbitMQConsumer) Close() {
	time.Sleep(Delay)

	if consumer.channel != nil && !consumer.channel.IsClosed() {
		log.Debug("rabbitmq consumer - closing connection")
		if err := consumer.channel.Close(); err != nil {
			log.Error(fmt.Sprintf("rabbitmq consumer - failed to close connection to queue %s: %s", consumer.name, err.Error()))
		}
	}
	consumer.channel = nil
	consumer.connection.Close()
	log.Debug(fmt.Sprintf("rabbitmq consumer - closed connection to queue %s", consumer.name))
}

func (consumer *RabbitMQConsumer) Context() Context {
	return consumer.connection.Context()
}
