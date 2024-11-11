package amqp

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

type consumer struct {
	connection       Connection
	listener         Listener
	channel          *amqp.Channel
	queue            amqp.Queue
	name             string
	consumer         string
	autoAck          bool
	noLocal          bool
	durable          bool
	autoDelete       bool
	exclusive        bool
	noWait           bool
	args             amqp.Table
	closingHandler   ClosingHandler
	messageProcessor MessageProcessor
	mu               sync.RWMutex
}

func NewConsumer(connection Connection, name string, options ...ConsumerOptions) *consumer {
	assert.NotNil(connection, "starting up - error setting up rabbitmq amqp consumer: connection is nil")
	assert.NotEmpty(name, "starting up - error setting up rabbitmq amqp consumer: name is empty")

	consumer := &consumer{
		connection:       connection,
		listener:         NewListener(),
		name:             name,
		consumer:         "consumer-" + name,
		autoAck:          false,
		noLocal:          false,
		durable:          false,
		autoDelete:       false,
		exclusive:        false,
		noWait:           false,
		args:             nil,
		closingHandler:   closingHandler,
		messageProcessor: messageProcessor,
	}

	for _, option := range options {
		option(consumer)
	}
	consumer.autoAck = false

	return consumer
}

func (consumer *consumer) Consume(ctx context.Context) (Event, error) {

	consumer.mu.Lock()
	defer consumer.mu.Unlock()

	var err error
	var connection *amqp.Connection
	if connection, err = consumer.connection.Connect(ctx); err != nil {
		log.Debug(ctx, fmt.Sprintf("rabbitmq consumer - failed connection to queue %s: %s", consumer.name, err.Error()))
		return nil, err
	}

	if !(consumer.channel != nil && !consumer.channel.IsClosed()) {
		if consumer.channel, err = connection.Channel(); err != nil {
			log.Debug(ctx, fmt.Sprintf("rabbitmq consumer - failed connection to queue %s: %s", consumer.name, err.Error()))
			return nil, err
		}
	}

	if consumer.queue, err = consumer.channel.QueueDeclare(consumer.name, consumer.durable, consumer.autoDelete, consumer.exclusive, consumer.noWait, consumer.args); err != nil {
		log.Debug(ctx, fmt.Sprintf("rabbitmq consumer - failed connection to queue %s: %s", consumer.name, err.Error()))
		return nil, err
	}

	log.Debug(ctx, fmt.Sprintf("rabbitmq consumer - connected to queue %s", consumer.name))

	var deliveries <-chan amqp.Delivery
	if deliveries, err = consumer.channel.ConsumeWithContext(ctx, consumer.name, consumer.consumer, consumer.autoAck, consumer.exclusive, consumer.noLocal, consumer.noWait, consumer.args); err != nil {
		log.Debug(ctx, fmt.Sprintf("rabbitmq consumer - failed comsuming from queue: %s", err.Error()))
		return nil, err
	}

	closeChannel := make(Event)
	go consumer.closingHandler(ctx, consumer.name, consumer.channel, deliveries, consumer.listener, closeChannel, consumer.messageProcessor)
	return closeChannel, nil
}

func (consumer *consumer) Close(ctx context.Context) {
	time.Sleep(Delay)

	if consumer.channel != nil && !consumer.channel.IsClosed() {
		log.Debug(ctx, "rabbitmq consumer - closing connection")
		if err := consumer.channel.Close(); err != nil {
			log.Error(ctx, fmt.Sprintf("rabbitmq consumer - failed to close connection to queue %s: %s", consumer.name, err.Error()))
		}
	}
	consumer.channel = nil
	consumer.connection.Close(ctx)
	log.Debug(ctx, fmt.Sprintf("rabbitmq consumer - closed connection to queue %s", consumer.name))
}

func (consumer *consumer) Context() Context {
	return consumer.connection.Context()
}

func (consumer *consumer) Set(property string, value any) {
	if utils.IsEmpty(property) || utils.IsEmpty(value) {
		return
	}

	switch property {
	case "listener":
		consumer.listener = utils.ToType[Listener](value)
	case "autoAck":
		consumer.autoAck = utils.ToBool(value)
	case "noLocal":
		consumer.noLocal = utils.ToBool(value)
	case "durable":
		consumer.durable = utils.ToBool(value)
	case "autoDelete":
		consumer.autoDelete = utils.ToBool(value)
	case "exclusive":
		consumer.exclusive = utils.ToBool(value)
	case "noWait":
		consumer.noWait = utils.ToBool(value)
	case "args":
		consumer.args = utils.ToType[amqp.Table](value)
	case "closingHandler":
		consumer.closingHandler = utils.ToType[ClosingHandler](value)
	case "messageProcessor":
		consumer.messageProcessor = utils.ToType[MessageProcessor](value)
	}
}
