package rabbitmq

import (
	"context"
	"fmt"
	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
	"github.com/guidomantilla/go-feather-lib/pkg/messaging"
	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/amqp"
	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/stream"
)

type StreamsConsumerOption func(*StreamsConsumer)

func WithStreamOptions(options *stream.StreamOptions) StreamsConsumerOption {
	return func(consumer *StreamsConsumer) {
		consumer.streamOptions = options
	}
}

func WithConsumerOptions(options *stream.ConsumerOptions) StreamsConsumerOption {
	return func(consumer *StreamsConsumer) {
		consumer.consumerOptions = options
	}
}

func WithRabbitMQStreamsListener(listener messaging.Listener[*amqp.Message]) StreamsConsumerOption {
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
