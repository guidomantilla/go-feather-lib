package rabbitmq

import (
	"context"
	"fmt"
	"sync"
	"time"

	samqp "github.com/rabbitmq/rabbitmq-stream-go-client/pkg/amqp"
	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/stream"

	"github.com/guidomantilla/go-feather-lib/pkg/common/assert"
	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
	"github.com/guidomantilla/go-feather-lib/pkg/messaging"
)

type StreamsProducer struct {
	connection      messaging.Connection[*stream.Environment]
	environment     *stream.Environment
	name            string
	streamOptions   *stream.StreamOptions
	producerOptions *stream.ProducerOptions
	mu              sync.RWMutex
}

func NewStreamsProducer(connection messaging.Connection[*stream.Environment], name string, options ...StreamsProducerOptions) *StreamsProducer {
	assert.NotNil(connection, "starting up - error setting up rabbitmq streams producer: connection is nil")
	assert.NotEmpty(name, "starting up - error setting up rabbitmq streams producer: name is empty")

	producer := &StreamsProducer{
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

func (streams *StreamsProducer) Produce(ctx context.Context, message *samqp.AMQP10) error {
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

func (streams *StreamsProducer) Close() {
	time.Sleep(messaging.Delay)

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

func (streams *StreamsProducer) Context() messaging.Context {
	return streams.connection.Context()
}
