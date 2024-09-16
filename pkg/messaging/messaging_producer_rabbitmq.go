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

type RabbitMQProducerOption func(*RabbitMQProducer)

func WithExchange(exchange string) RabbitMQProducerOption {
	return func(producer *RabbitMQProducer) {
		producer.exchange = exchange
	}
}

func WithMandatory(mandatory bool) RabbitMQProducerOption {
	return func(producer *RabbitMQProducer) {
		producer.mandatory = mandatory
	}
}

func WithImmediate(immediate bool) RabbitMQProducerOption {
	return func(producer *RabbitMQProducer) {
		producer.immediate = immediate
	}
}

type RabbitMQProducer struct {
	messagingConnection MessagingConnection[*amqp.Connection]
	channel             *amqp.Channel
	name                string
	exchange            string
	mandatory           bool
	immediate           bool
	mu                  sync.Mutex
}

func NewRabbitMQProducer(messagingConnection MessagingConnection[*amqp.Connection], name string, options ...RabbitMQProducerOption) *RabbitMQProducer {

	if messagingConnection == nil {
		log.Fatal("starting up - error setting up rabbitmq producer: messagingConnection is nil")
	}

	if strings.TrimSpace(name) == "" {
		log.Fatal("starting up - error setting up rabbitmq producer: name is empty")
	}

	producer := &RabbitMQProducer{
		messagingConnection: messagingConnection,
		name:                name,
		exchange:            "",
		mandatory:           false,
		immediate:           false,
	}

	for _, option := range options {
		option(producer)
	}

	return producer
}

func (producer *RabbitMQProducer) Produce(ctx context.Context, message *amqp.Publishing) error {
	producer.mu.Lock()
	defer producer.mu.Unlock()

	var err error
	var connection *amqp.Connection
	if connection, err = producer.messagingConnection.Connect(); err != nil {
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

func (producer *RabbitMQProducer) Close() {
	time.Sleep(MessagingDelay)

	if producer.channel != nil && !producer.channel.IsClosed() {
		log.Debug("rabbitmq producer - closing connection")
		if err := producer.channel.Close(); err != nil {
			log.Error(fmt.Sprintf("rabbitmq producer - failed to close connection to queue %s: %s", producer.name, err.Error()))
		}
	}
	producer.channel = nil
	producer.messagingConnection.Close()
	log.Debug(fmt.Sprintf("rabbitmq producer - closed connection to queue %s", producer.name))
}

func (producer *RabbitMQProducer) MessagingContext() MessagingContext {
	return producer.messagingConnection.MessagingContext()
}
