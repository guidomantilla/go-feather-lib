package messaging

import (
	"context"
	"fmt"
	"strings"
	"sync"

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

func WithRabbitMQListener(listener MessagingListener[*amqp.Delivery]) RabbitMQConsumerOption {
	return func(consumer *RabbitMQConsumer) {
		consumer.listener = listener
	}
}

type RabbitMQConsumer struct {
	messagingConnection MessagingConnection[*amqp.Connection]
	listener            MessagingListener[*amqp.Delivery]
	channel             *amqp.Channel
	queue               amqp.Queue
	name                string
	consumer            string
	autoAck             bool
	noLocal             bool
	durable             bool
	autoDelete          bool
	exclusive           bool
	noWait              bool
	args                amqp.Table
	mu                  sync.Mutex
}

func NewRabbitMQConsumer(messagingConnection MessagingConnection[*amqp.Connection], name string, options ...RabbitMQConsumerOption) *RabbitMQConsumer {

	if messagingConnection == nil {
		log.Fatal("starting up - error setting up rabbitmq consumer: messagingConnection is nil")
	}

	if strings.TrimSpace(name) == "" {
		log.Fatal("starting up - error setting up rabbitmq consumer: name is empty")
	}

	consumer := &RabbitMQConsumer{
		messagingConnection: messagingConnection,
		listener:            NewRabbitMQListener(),
		name:                name,
		consumer:            "consumer-" + name,
		autoAck:             true,
		noLocal:             false,
		durable:             false,
		autoDelete:          false,
		exclusive:           false,
		noWait:              false,
		args:                nil,
	}

	for _, option := range options {
		option(consumer)
	}

	return consumer
}

func (queue *RabbitMQConsumer) Consume(ctx context.Context) (MessagingEvent, error) {

	queue.mu.Lock()
	defer queue.mu.Unlock()

	var err error
	var connection *amqp.Connection
	if connection, err = queue.messagingConnection.Connect(); err != nil {
		log.Debug(fmt.Sprintf("rabbitmq consumer - failed connection to queue %s: %s", queue.name, err.Error()))
		return nil, err
	}

	if !(queue.channel != nil && !queue.channel.IsClosed()) {
		if queue.channel, err = connection.Channel(); err != nil {
			log.Debug(fmt.Sprintf("rabbitmq consumer - failed connection to queue %s: %s", queue.name, err.Error()))
			return nil, err
		}
	}

	if queue.queue, err = queue.channel.QueueDeclare(queue.name, queue.durable, queue.autoDelete, queue.exclusive, queue.noWait, queue.args); err != nil {
		log.Debug(fmt.Sprintf("rabbitmq consumer - failed connection to queue %s: %s", queue.name, err.Error()))
		return nil, err
	}

	log.Debug(fmt.Sprintf("rabbitmq consumer - connected to queue %s", queue.name))

	var deliveries <-chan amqp.Delivery
	if deliveries, err = queue.channel.ConsumeWithContext(ctx, queue.name, queue.consumer, queue.autoAck, queue.exclusive, queue.noLocal, queue.noWait, queue.args); err != nil {
		log.Debug(fmt.Sprintf("rabbitmq consumer - failed comsuming from queue: %s", err.Error()))
		return nil, err
	}

	closeChannel := make(chan string)
	closeHandler := func(ctx context.Context, listener MessagingListener[*amqp.Delivery], channel *amqp.Channel, queue string, closeChannel chan string) {
		var err error
		for message := range deliveries {
			go func() {
				log.Debug(fmt.Sprintf("rabbitmq consumer - message received: %s", message.Body))
				if err := listener.OnMessage(ctx, &message); err != nil {
					log.Debug(fmt.Sprintf("rabbitmq consumer - failed to process message: %s", err.Error()))
				}
			}()
		}
		if err = channel.Close(); err != nil {
			log.Debug(fmt.Sprintf("rabbitmq consumer - failed to close channel to queue %s: %s", queue, err.Error()))
			return
		}
		close(closeChannel)
		log.Debug(fmt.Sprintf("rabbitmq consumer - disconected from queue %s", queue))
	}

	go closeHandler(ctx, queue.listener, queue.channel, queue.name, closeChannel)
	return closeChannel, nil
}

func (queue *RabbitMQConsumer) Close() {
	if queue.channel != nil && !queue.channel.IsClosed() {
		log.Debug("rabbitmq consumer - closing connection")
		if err := queue.channel.Close(); err != nil {
			log.Error(fmt.Sprintf("rabbitmq consumer - failed to close connection to queue %s: %s", queue.name, err.Error()))
		}
	}
	queue.channel = nil
	queue.messagingConnection.Close()
	log.Debug(fmt.Sprintf("rabbitmq consumer - closed connection to queue %s", queue.name))
}

func (queue *RabbitMQConsumer) MessagingContext() MessagingContext {
	return queue.messagingConnection.MessagingContext()
}
