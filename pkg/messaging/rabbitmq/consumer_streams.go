package rabbitmq

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/amqp"
	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/stream"

	"github.com/guidomantilla/go-feather-lib/pkg/common/assert"
	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type StreamsConsumer struct {
	connection      Connection[*stream.Environment]
	listener        Listener[*amqp.Message]
	environment     *stream.Environment
	name            string
	consumer        string
	streamOptions   *stream.StreamOptions
	consumerOptions *stream.ConsumerOptions
	messagesHandler stream.MessagesHandler
	mu              sync.RWMutex
}

func NewStreamsConsumer(connection Connection[*stream.Environment], name string, options ...StreamsConsumerOptions) *StreamsConsumer {
	assert.NotNil(connection, "starting up - error setting up rabbitmq streams consumer: connection is nil")
	assert.NotEmpty(name, "starting up - error setting up rabbitmq streams consumer: name is empty")

	listener := NewStreamsListener()
	consumer := &StreamsConsumer{
		connection:      connection,
		name:            name,
		consumer:        "consumer-" + name,
		streamOptions:   stream.NewStreamOptions(),
		consumerOptions: stream.NewConsumerOptions().SetConsumerName("consumer-" + name),
		listener:        listener,
		messagesHandler: func(consumerContext stream.ConsumerContext, message *amqp.Message) {
			go func(consumerContext stream.ConsumerContext, message *amqp.Message) {
				log.Debug(fmt.Sprintf("rabbitmq streams consumer - message received: %s", message.Data))
				ctx := context.WithValue(context.Background(), stream.ConsumerContext{}, consumerContext)
				if err := listener.OnMessage(ctx, message); err != nil {
					log.Debug(fmt.Sprintf("rabbitmq streams consumer - failed to process message: %s", err.Error()))
					return
				}
			}(consumerContext, message)
		},
	}

	for _, option := range options {
		option(consumer)
	}

	consumer.consumerOptions.SetConsumerName(consumer.consumer)
	consumer.consumerOptions.SetManualCommit()

	return consumer
}

func (streams *StreamsConsumer) Consume(ctx context.Context) (Event, error) {
	streams.mu.Lock()
	defer streams.mu.Unlock()

	var err error
	if streams.environment, err = streams.connection.Connect(); err != nil {
		log.Debug(fmt.Sprintf("rabbitmq streams consumer - failed connection to stream %s: %s", streams.name, err.Error()))
		return nil, err
	}

	var streamExists bool
	if streamExists, err = streams.environment.StreamExists(streams.name); err != nil {
		log.Debug(fmt.Sprintf("rabbitmq streams consumer - failed connection to stream %s: %s", streams.name, err.Error()))
		return nil, err
	}

	if !streamExists {
		if err = streams.environment.DeclareStream(streams.name, streams.streamOptions); err != nil {
			log.Debug(fmt.Sprintf("rabbitmq streams consumer - failed connection to stream %s: %s", streams.name, err.Error()))
			return nil, err
		}
	}

	log.Debug(fmt.Sprintf("rabbitmq streams consumer - connected to stream %s", streams.name))

	var storedOffset int64
	if storedOffset, err = streams.environment.QueryOffset(streams.consumer, streams.name); err != nil {
		if errors.Is(err, stream.OffsetNotFoundError) {
			log.Debug(fmt.Sprintf("rabbitmq streams consumer - failed to query offset from stream %s: %s", streams.name, err.Error()))
			log.Debug(fmt.Sprintf("rabbitmq streams consumer - setting up offset to FIRST from stream %s", streams.name))
			streams.consumerOptions.SetOffset(stream.OffsetSpecification{}.First())
		} else {
			newOffset := storedOffset + 1
			log.Debug(fmt.Sprintf("rabbitmq streams consumer - setting up offset to %d from stream %s", newOffset, streams.name))
			streams.consumerOptions.SetOffset(stream.OffsetSpecification{}.Offset(newOffset))
		}
	}

	var consumer *stream.Consumer
	if consumer, err = streams.environment.NewConsumer(streams.name, streams.messagesHandler, streams.consumerOptions); err != nil {
		log.Debug(fmt.Sprintf("rabbitmq streams consumer - failed comsuming from stream %s: %s", streams.name, err.Error()))
		return nil, err
	}

	closeChannel := make(chan string)
	closeHandler := func(consumer *stream.Consumer, stream string, closeChannel chan string) {
		var err error
		for range consumer.NotifyClose() {
			if err := consumer.StoreOffset(); err != nil {
				log.Debug(fmt.Sprintf("rabbitmq streams consumer - failed to store consumer offset from stream %s: %s", stream, err.Error()))
				return
			}
			if err = consumer.Close(); err != nil {
				log.Debug(fmt.Sprintf("rabbitmq streams consumer - failed to close consumer from stream %s: %s", stream, err.Error()))
				return
			}
			close(closeChannel)
		}
		log.Debug(fmt.Sprintf("rabbitmq streams consumer - disconected from stream %s", stream))
	}

	go closeHandler(consumer, streams.name, closeChannel)
	return closeChannel, nil
}

func (streams *StreamsConsumer) Close() {
	time.Sleep(Delay)

	if streams.environment != nil && !streams.environment.IsClosed() {
		log.Debug("rabbitmq streams consumer - closing connection")
		if err := streams.environment.Close(); err != nil {
			log.Error(fmt.Sprintf("rabbitmq streams consumer - failed to close connection to stream %s: %s", streams.name, err.Error()))
		}
	}
	streams.environment = nil
	streams.connection.Close()
	log.Debug(fmt.Sprintf("rabbitmq streams consumer - closed connection to stream %s", streams.name))
}

func (streams *StreamsConsumer) Context() Context {
	return streams.connection.Context()
}

func (streams *StreamsConsumer) Set(property string, value any) {

}
