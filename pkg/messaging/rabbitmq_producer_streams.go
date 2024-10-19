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
	connection      Connection[*stream.Environment]
	environment     *stream.Environment
	name            string
	streamOptions   *stream.StreamOptions
	producerOptions *stream.ProducerOptions
	mu              sync.RWMutex
}

func NewRabbitMQStreamsProducer(connection Connection[*stream.Environment], name string, options ...RabbitMQStreamsProducerOption) *RabbitMQStreamsProducer {

	if connection == nil {
		log.Fatal("starting up - error setting up rabbitmq streams producer: connection is nil")
	}

	if strings.TrimSpace(name) == "" {
		log.Fatal("starting up - error setting up rabbitmq streams producer: name is empty")
	}

	producer := &RabbitMQStreamsProducer{
		connection:      connection,
		name:            name,
		streamOptions:   stream.NewStreamOptions(),
		producerOptions: stream.NewProducerOptions(),
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
	if streams.environment, err = streams.connection.Connect(); err != nil {
		log.Debug(fmt.Sprintf("rabbitmq streams producer - failed connection to stream %s: %s", streams.name, err.Error()))
		return err
	}

	var streamExists bool
	if streamExists, err = streams.environment.StreamExists(streams.name); err != nil {
		log.Debug(fmt.Sprintf("rabbitmq streams producer - failed connection to stream %s: %s", streams.name, err.Error()))
		return err
	}

	if !streamExists {
		if err = streams.environment.DeclareStream(streams.name, streams.streamOptions); err != nil {
			log.Debug(fmt.Sprintf("rabbitmq streams consumer - failed connection to stream %s: %s", streams.name, err.Error()))
			return err
		}
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
	time.Sleep(Delay)

	if streams.environment != nil && !streams.environment.IsClosed() {
		log.Debug("rabbitmq streams producer - closing connection")
		if err := streams.environment.Close(); err != nil {
			log.Error(fmt.Sprintf("rabbitmq streams producer - failed to close connection to stream %s: %s", streams.name, err.Error()))
		}
	}
	streams.environment = nil
	streams.connection.Close()
	log.Debug(fmt.Sprintf("rabbitmq streams consumer - producer connection to stream %s", streams.name))
}

func (streams *RabbitMQStreamsProducer) Context() Context {
	return streams.connection.Context()
}
