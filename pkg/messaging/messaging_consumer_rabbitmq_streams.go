package messaging

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/amqp"
	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/stream"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type RabbitMQStreamsConsumerOption func(*RabbitMQStreamsConsumer)

func WithStreamOptions(options *stream.StreamOptions) RabbitMQStreamsConsumerOption {
	return func(consumer *RabbitMQStreamsConsumer) {
		consumer.streamOptions = options
	}
}

func WithConsumerOptions(options *stream.ConsumerOptions) RabbitMQStreamsConsumerOption {
	return func(consumer *RabbitMQStreamsConsumer) {
		consumer.consumerOptions = options
	}
}

func WithRabbitMQStreamsListener(listener MessagingListener[*amqp.Message]) RabbitMQStreamsConsumerOption {
	return func(consumer *RabbitMQStreamsConsumer) {
		consumer.listener = listener
		consumer.messagesHandler = func(consumerContext stream.ConsumerContext, message *amqp.Message) {
			log.Debug(fmt.Sprintf("rabbitmq streams consumer - message received: %s", message.Data))
			if err := listener.OnMessage(context.Background(), message); err != nil {
				log.Debug(fmt.Sprintf("rabbitmq streams consumer - failed to process message: %s", err.Error()))
			}
		}
	}
}

type RabbitMQStreamsConsumer struct {
	messagingConnection MessagingConnection[*stream.Environment]
	listener            MessagingListener[*amqp.Message]
	environment         *stream.Environment
	name                string
	consumer            string
	streamOptions       *stream.StreamOptions
	consumerOptions     *stream.ConsumerOptions
	messagesHandler     stream.MessagesHandler
	mu                  sync.RWMutex
}

func NewRabbitMQStreamsConsumer(messagingConnection MessagingConnection[*stream.Environment], name string, options ...RabbitMQStreamsConsumerOption) *RabbitMQStreamsConsumer {

	if messagingConnection == nil {
		log.Fatal("starting up - error setting up rabbitmq streams consumer: messagingConnection is nil")
	}

	if strings.TrimSpace(name) == "" {
		log.Fatal("starting up - error setting up rabbitmq streams consumer: name is empty")
	}
	listener := NewRabbitMQStreamsListener()
	consumer := &RabbitMQStreamsConsumer{
		messagingConnection: messagingConnection,
		name:                name,
		consumer:            "consumer-" + name,
		streamOptions:       stream.NewStreamOptions(),
		consumerOptions:     stream.NewConsumerOptions().SetConsumerName("consumer-" + name),
		listener:            listener,
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

func (streams *RabbitMQStreamsConsumer) Consume(ctx context.Context) (MessagingEvent, error) {
	streams.mu.Lock()
	defer streams.mu.Unlock()

	var err error
	if streams.environment, err = streams.messagingConnection.Connect(); err != nil {
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

func (streams *RabbitMQStreamsConsumer) Close() {
	time.Sleep(Delay)

	if streams.environment != nil && !streams.environment.IsClosed() {
		log.Debug("rabbitmq streams consumer - closing connection")
		if err := streams.environment.Close(); err != nil {
			log.Error(fmt.Sprintf("rabbitmq streams consumer - failed to close connection to stream %s: %s", streams.name, err.Error()))
		}
	}
	streams.environment = nil
	streams.messagingConnection.Close()
	log.Debug(fmt.Sprintf("rabbitmq streams consumer - closed connection to stream %s", streams.name))
}

func (streams *RabbitMQStreamsConsumer) MessagingContext() MessagingContext {
	return streams.messagingConnection.MessagingContext()
}
