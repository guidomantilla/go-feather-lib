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

type producer struct {
	connection Connection
	channel    *amqp.Channel
	name       string
	exchange   string
	mandatory  bool
	immediate  bool
	mu         sync.RWMutex
}

func NewProducer(connection Connection, name string, options ...ProducerOptions) *producer {
	assert.NotNil(connection, "starting up - error setting up rabbitmq amqp producer: connection is nil")
	assert.NotEmpty(name, "starting up - error setting up rabbitmq amqp producer: name is empty")

	producer := &producer{
		connection: connection,
		name:       name,
		exchange:   "",
		mandatory:  false,
		immediate:  false,
	}

	for _, option := range options {
		option(producer)
	}

	return producer
}

func (producer *producer) Produce(ctx context.Context, message *amqp.Publishing) error {
	producer.mu.Lock()
	defer producer.mu.Unlock()

	var err error
	var connection *amqp.Connection
	if connection, err = producer.connection.Connect(ctx); err != nil {
		log.Debug(ctx, fmt.Sprintf("rabbitmq producer - failed connection to queue %s: %s", producer.name, err.Error()))
		return err
	}

	if !(producer.channel != nil && !producer.channel.IsClosed()) {
		if producer.channel, err = connection.Channel(); err != nil {
			log.Debug(ctx, fmt.Sprintf("rabbitmq producer - failed connection to queue %s: %s", producer.name, err.Error()))
			return err
		}
	}

	log.Debug(ctx, fmt.Sprintf("rabbitmq producer - publishing to queue %s", producer.name))
	if err = producer.channel.PublishWithContext(ctx, producer.exchange, producer.name, producer.mandatory, producer.immediate, *message); err != nil {
		log.Debug(ctx, fmt.Sprintf("rabbitmq producer - failed publishing to queue: %s", err.Error()))
		return err
	}
	log.Debug(ctx, fmt.Sprintf("rabbitmq producer - published to queue %s", producer.name))
	return nil
}

func (producer *producer) Close(ctx context.Context) {
	time.Sleep(Delay)

	if producer.channel != nil && !producer.channel.IsClosed() {
		log.Debug(ctx, "rabbitmq producer - closing connection")
		if err := producer.channel.Close(); err != nil {
			log.Error(ctx, fmt.Sprintf("rabbitmq producer - failed to close connection to queue %s: %s", producer.name, err.Error()))
		}
	}
	producer.channel = nil
	producer.connection.Close(ctx)
	log.Debug(ctx, fmt.Sprintf("rabbitmq producer - closed connection to queue %s", producer.name))
}

func (producer *producer) Context() Context {
	return producer.connection.Context()
}

func (producer *producer) Set(property string, value any) {
	if utils.IsEmpty(property) || utils.IsEmpty(value) {
		return
	}

	switch property {
	case "exchange":
		producer.exchange = utils.ToString(value)
	case "mandatory":
		producer.mandatory = utils.ToBool(value)
	case "immediate":
		producer.immediate = utils.ToBool(value)
	}
}
