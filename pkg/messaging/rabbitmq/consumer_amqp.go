package rabbitmq

import (
	"context"
	"fmt"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/guidomantilla/go-feather-lib/pkg/common/assert"
	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
	"github.com/guidomantilla/go-feather-lib/pkg/common/utils"
)

type amqpConsumer struct {
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

func NewAmqpConsumer(connection Connection[*amqp.Connection], name string, options ...AmqpConsumerOptions) *amqpConsumer {
	assert.NotNil(connection, "starting up - error setting up rabbitmq amqp consumer: connection is nil")
	assert.NotEmpty(name, "starting up - error setting up rabbitmq amqp consumer: name is empty")

	consumer := &amqpConsumer{
		connection: connection,
		listener:   NewAmqpListener(),
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

func (consumer *amqpConsumer) Consume(ctx context.Context) (Event, error) {

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

func (consumer *amqpConsumer) Close() {
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

func (consumer *amqpConsumer) Context() Context {
	return consumer.connection.Context()
}

func (consumer *amqpConsumer) Set(property string, value any) {
	if utils.IsEmpty(property) || utils.IsEmpty(value) {
		return
	}

	switch property {
	case "listener":
		if value != nil {
			consumer.listener = value.(Listener[*amqp.Delivery])
		}
	case "autoAck":
		consumer.autoAck = value.(bool)
	case "noLocal":
		consumer.noLocal = value.(bool)
	case "durable":
		consumer.durable = value.(bool)
	case "autoDelete":
		consumer.autoDelete = value.(bool)
	case "exclusive":
		consumer.exclusive = value.(bool)
	case "noWait":
		consumer.noWait = value.(bool)
	case "args":
		if value != nil {
			consumer.args = value.(amqp.Table)
		}
	}
}
