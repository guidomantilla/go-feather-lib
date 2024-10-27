package streams

import (
	"context"
	"fmt"

	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/amqp"
	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/stream"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

var consumerOptions = NewConsumerOptions()

func NewConsumerOptions() ConsumerOptions {
	return func(consumer *consumer) {
	}
}

type ConsumerOptions func(*consumer)

func (options ConsumerOptions) WithStreamOptions(soptions *stream.StreamOptions) ConsumerOptions {
	return func(consumer *consumer) {
		consumer.streamOptions = soptions
	}
}

func (options ConsumerOptions) WithConsumerOptions(coptions *stream.ConsumerOptions) ConsumerOptions {
	return func(consumer *consumer) {
		consumer.consumerOptions = coptions
	}
}

func (options ConsumerOptions) WithListener(listener Listener) ConsumerOptions {
	return func(consumer *consumer) {
		consumer.listener = listener
		consumer.messagesHandler = func(consumerContext stream.ConsumerContext, message *amqp.Message) {
			log.Debug(fmt.Sprintf("rabbitmq streams consumer - message received: %s", message.Data))
			if err := listener.OnMessage(context.Background(), message); err != nil {
				log.Debug(fmt.Sprintf("rabbitmq streams consumer - failed to process message: %s", err.Error()))
			}
		}
	}
}
