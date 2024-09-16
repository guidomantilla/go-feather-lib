package messaging

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	samqp "github.com/rabbitmq/rabbitmq-stream-go-client/pkg/amqp"
	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/stream"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type RabbitMQStreamsProducerOption func(*RabbitMQStreamsProducer)

func WithProducerOptions(options *stream.ProducerOptions) RabbitMQStreamsProducerOption {
	return func(producer *RabbitMQStreamsProducer) {
		producer.producerOptions = options
	}
}

type RabbitMQStreamsProducer struct {
	messagingConnection MessagingConnection[*stream.Environment]
	environment         *stream.Environment
	name                string
	producerOptions     *stream.ProducerOptions
	mu                  sync.Mutex
}

func NewRabbitMQStreamsProducer(messagingConnection MessagingConnection[*stream.Environment], name string, options ...RabbitMQStreamsProducerOption) *RabbitMQStreamsProducer {

	if messagingConnection == nil {
		log.Fatal("starting up - error setting up rabbitmq streams producer: messagingConnection is nil")
	}

	if strings.TrimSpace(name) == "" {
		log.Fatal("starting up - error setting up rabbitmq streams producer: name is empty")
	}

	producer := &RabbitMQStreamsProducer{
		messagingConnection: messagingConnection,
		name:                name,
		producerOptions:     stream.NewProducerOptions(),
	}

	for _, option := range options {
		option(producer)
	}

	return producer
}

func (streams *RabbitMQStreamsProducer) Produce(ctx context.Context, message *samqp.AMQP10) error {
	streams.mu.Lock()
	defer streams.mu.Unlock()

	var err error
	if streams.environment, err = streams.messagingConnection.Connect(); err != nil {
		log.Debug(fmt.Sprintf("rabbitmq streams producer - failed connection to stream %s: %s", streams.name, err.Error()))
		return err
	}

	var producer *stream.Producer
	if producer, err = streams.environment.NewProducer(streams.name, streams.producerOptions); err != nil {
		log.Debug(fmt.Sprintf("rabbitmq streams producer - failed connection to stream %s: %s", streams.name, err.Error()))
		return err
	}

	log.Debug(fmt.Sprintf("rabbitmq producer - publishing to stream %s", streams.name))
	if err = producer.Send(message); err != nil {
		log.Debug(fmt.Sprintf("rabbitmq streams producer - failed publishing message to stream %s: %s", streams.name, err.Error()))
		return err
	}
	log.Debug(fmt.Sprintf("rabbitmq producer - published to stream %s", streams.name))
	return nil
}

func (streams *RabbitMQStreamsProducer) Close() {
	time.Sleep(MessagingDelay)

	if streams.environment != nil && !streams.environment.IsClosed() {
		log.Debug("rabbitmq streams producer - closing connection")
		if err := streams.environment.Close(); err != nil {
			log.Error(fmt.Sprintf("rabbitmq streams producer - failed to close connection to stream %s: %s", streams.name, err.Error()))
		}
	}
	streams.environment = nil
	streams.messagingConnection.Close()
	log.Debug(fmt.Sprintf("rabbitmq streams consumer - producer connection to stream %s", streams.name))
}

func (streams *RabbitMQStreamsProducer) MessagingContext() MessagingContext {
	return streams.messagingConnection.MessagingContext()
}
