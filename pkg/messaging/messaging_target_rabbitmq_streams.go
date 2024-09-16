package messaging

import (
	"fmt"
	"strings"
	"sync"

	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/amqp"
	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/stream"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type RabbitMQStreams struct {
	messagingConnection MessagingConnection[*stream.Environment]
	environment         *stream.Environment
	name                string
	consumer            string
	mu                  sync.Mutex
}

func NewRabbitMQStreams(messagingConnection MessagingConnection[*stream.Environment], stream string) *RabbitMQStreams {

	if messagingConnection == nil {
		log.Fatal("starting up - error setting up rabbitMQStreams: messagingConnection is nil")
	}

	if strings.TrimSpace(stream) == "" {
		log.Fatal("starting up - error setting up rabbitMQStreams: stream is empty")
	}

	return &RabbitMQStreams{
		messagingConnection: messagingConnection,
		name:                stream,
		consumer:            "consumer-" + stream,
	}
}

func (streams *RabbitMQStreams) Consume() (MessagingEvent, error) {

	streams.mu.Lock()
	defer streams.mu.Unlock()

	var err error
	if streams.environment, err = streams.messagingConnection.Connect(); err != nil {
		log.Debug(fmt.Sprintf("rabbitmq streams - failed connection to stream %s: %s", streams.name, err.Error()))
		return nil, err
	}

	streamOptions := stream.NewStreamOptions().SetMaxLengthBytes(stream.ByteCapacity{}.GB(2))
	if err = streams.environment.DeclareStream(streams.name, streamOptions); err != nil {
		log.Debug(fmt.Sprintf("rabbitmq streams - failed connection to stream %s: %s", streams.name, err.Error()))
		return nil, err
	}

	log.Debug(fmt.Sprintf("rabbitmq streams - connected to stream %s", streams.name))

	messagesHandler := func(consumerContext stream.ConsumerContext, message *amqp.Message) {
		log.Info(fmt.Sprintf("rabbitmq queue - message received: %s", message.Data))
	}

	var consumer *stream.Consumer
	consumerOptions := stream.NewConsumerOptions().SetOffset(stream.OffsetSpecification{}.First()).SetConsumerName(streams.consumer)
	if consumer, err = streams.environment.NewConsumer(streams.name, messagesHandler, consumerOptions); err != nil {
		log.Debug(fmt.Sprintf("rabbitmq streams - failed comsuming from stream %s: %s", streams.name, err.Error()))
		return nil, err
	}

	closeChannel := make(chan string)
	go func(closeChannel chan string) {
		for range consumer.NotifyClose() {
			close(closeChannel)
			break
		}
		log.Debug(fmt.Sprintf("rabbitmq streams - disconected to queue %s", streams.name))
	}(closeChannel)

	return closeChannel, nil
}

func (streams *RabbitMQStreams) Close() {
	if streams.environment != nil && !streams.environment.IsClosed() {
		log.Debug("rabbitmq streams - closing connection")
		if err := streams.environment.Close(); err != nil {
			log.Error(fmt.Sprintf("rabbitmq streams - failed to close connection to stream %s: %s", streams.name, err.Error()))
		}
	}
	streams.environment = nil
	streams.messagingConnection.Close()
	log.Debug(fmt.Sprintf("rabbitmq streams - closed connection to stream %s", streams.name))
}

func (streams *RabbitMQStreams) MessagingContext() MessagingContext {
	return streams.messagingConnection.MessagingContext()
}
