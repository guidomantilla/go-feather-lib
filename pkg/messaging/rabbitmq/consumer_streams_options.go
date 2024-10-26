package rabbitmq

import (
	"context"
	"fmt"

	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/amqp"
	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/stream"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

var streamsConsumerOptions = NewStreamsConsumerOptions()

func NewStreamsConsumerOptions() StreamsConsumerOptions {
	return func(consumer *StreamsConsumer) {
	}
}

type StreamsConsumerOptions func(*StreamsConsumer)

func (options StreamsConsumerOptions) WithStreamOptions(soptions *stream.StreamOptions) StreamsConsumerOptions {
	return func(consumer *StreamsConsumer) {
		consumer.streamOptions = soptions
	}
}

func (options StreamsConsumerOptions) WithConsumerOptions(coptions *stream.ConsumerOptions) StreamsConsumerOptions {
	return func(consumer *StreamsConsumer) {
		consumer.consumerOptions = coptions
	}
}

func (options StreamsConsumerOptions) WithStreamsListener(listener Listener[*amqp.Message]) StreamsConsumerOptions {
	return func(consumer *StreamsConsumer) {
		consumer.listener = listener
		consumer.messagesHandler = func(consumerContext stream.ConsumerContext, message *amqp.Message) {
			log.Debug(fmt.Sprintf("rabbitmq streams consumer - message received: %s", message.Data))
			if err := listener.OnMessage(context.Background(), message); err != nil {
				log.Debug(fmt.Sprintf("rabbitmq streams consumer - failed to process message: %s", err.Error()))
			}
		}
	}
}
