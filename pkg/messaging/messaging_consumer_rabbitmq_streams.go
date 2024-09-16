package messaging

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/amqp"
	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/stream"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type RabbitMQStreamsConsumer struct {
	messagingConnection MessagingConnection[*stream.Environment]
	environment         *stream.Environment
	name                string
	consumer            string
	mu                  sync.Mutex
}

func NewRabbitMQStreamsConsumer(messagingConnection MessagingConnection[*stream.Environment], stream string) *RabbitMQStreamsConsumer {

	if messagingConnection == nil {
		log.Fatal("starting up - error setting up rabbitmq streams consumer: messagingConnection is nil")
	}

	if strings.TrimSpace(stream) == "" {
		log.Fatal("starting up - error setting up rabbitmq streams consumer: stream is empty")
	}

	return &RabbitMQStreamsConsumer{
		messagingConnection: messagingConnection,
		name:                stream,
		consumer:            "consumer-" + stream,
	}
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
		streamOptions := stream.NewStreamOptions().SetMaxLengthBytes(stream.ByteCapacity{}.GB(2))
		if err = streams.environment.DeclareStream(streams.name, streamOptions); err != nil {
			log.Debug(fmt.Sprintf("rabbitmq streams consumer - failed connection to stream %s: %s", streams.name, err.Error()))
			return nil, err
		}
	}

	log.Debug(fmt.Sprintf("rabbitmq streams consumer - connected to stream %s", streams.name))

	messagesHandler := func(consumerContext stream.ConsumerContext, message *amqp.Message) {
		log.Info(fmt.Sprintf("rabbitmq streams consumer - message received: %s", message.Data))
	}

	var consumer *stream.Consumer
	consumerOptions := stream.NewConsumerOptions().SetOffset(stream.OffsetSpecification{}.First()).SetConsumerName(streams.consumer)
	if consumer, err = streams.environment.NewConsumer(streams.name, messagesHandler, consumerOptions); err != nil {
		log.Debug(fmt.Sprintf("rabbitmq streams consumer - failed comsuming from stream %s: %s", streams.name, err.Error()))
		return nil, err
	}

	closeChannel := make(chan string)
	closeHandler := func(consumer *stream.Consumer, stream string, closeChannel chan string) {
		var err error
		for range consumer.NotifyClose() {
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
	if streams.environment != nil && !streams.environment.IsClosed() {
		log.Debug("rabbitmq streams - closing connection")
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
