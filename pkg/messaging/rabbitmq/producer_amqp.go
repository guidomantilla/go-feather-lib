package rabbitmq

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
	"github.com/guidomantilla/go-feather-lib/pkg/messaging"
)

type AmqpProducer struct {
	connection messaging.Connection[*amqp.Connection]
	channel    *amqp.Channel
	name       string
	exchange   string
	mandatory  bool
	immediate  bool
	mu         sync.RWMutex
}

func NewAmqpProducer(connection messaging.Connection[*amqp.Connection], name string, options ...AmqpProducerOptions) *AmqpProducer {

	if connection == nil {
		log.Fatal("starting up - error setting up rabbitmq producer: connection is nil")
	}

	if strings.TrimSpace(name) == "" {
		log.Fatal("starting up - error setting up rabbitmq producer: name is empty")
	}

	producer := &AmqpProducer{
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

func (producer *AmqpProducer) Produce(ctx context.Context, message *amqp.Publishing) error {
	producer.mu.Lock()
	defer producer.mu.Unlock()

	var err error
	var connection *amqp.Connection
	if connection, err = producer.connection.Connect(); err != nil {
		log.Debug(fmt.Sprintf("rabbitmq producer - failed connection to queue %s: %s", producer.name, err.Error()))
		return err
	}

	if !(producer.channel != nil && !producer.channel.IsClosed()) {
		if producer.channel, err = connection.Channel(); err != nil {
			log.Debug(fmt.Sprintf("rabbitmq producer - failed connection to queue %s: %s", producer.name, err.Error()))
			return err
		}
	}

	log.Debug(fmt.Sprintf("rabbitmq producer - publishing to queue %s", producer.name))
	if err = producer.channel.PublishWithContext(ctx, producer.exchange, producer.name, producer.mandatory, producer.immediate, *message); err != nil {
		log.Debug(fmt.Sprintf("rabbitmq producer - failed publishing to queue: %s", err.Error()))
		return err
	}
	log.Debug(fmt.Sprintf("rabbitmq producer - published to queue %s", producer.name))
	return nil
}

func (producer *AmqpProducer) Close() {
	time.Sleep(messaging.Delay)

	if producer.channel != nil && !producer.channel.IsClosed() {
		log.Debug("rabbitmq producer - closing connection")
		if err := producer.channel.Close(); err != nil {
			log.Error(fmt.Sprintf("rabbitmq producer - failed to close connection to queue %s: %s", producer.name, err.Error()))
		}
	}
	producer.channel = nil
	producer.connection.Close()
	log.Debug(fmt.Sprintf("rabbitmq producer - closed connection to queue %s", producer.name))
}

func (producer *AmqpProducer) Context() messaging.Context {
	return producer.connection.Context()
}
