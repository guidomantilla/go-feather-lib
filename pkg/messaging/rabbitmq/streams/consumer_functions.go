package streams

import (
	"context"
	"fmt"

	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/amqp"
	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/stream"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type ClosingHandler func(ctx context.Context, consumer *stream.Consumer, stream string, closeChannel chan string)

func closingHandler(ctx context.Context, consumer *stream.Consumer, stream string, closeChannel chan string) {
	var err error
	for range consumer.NotifyClose() {
		if err := consumer.StoreOffset(); err != nil {
			log.Debug(ctx, fmt.Sprintf("rabbitmq streams consumer - failed to store consumer offset from stream %s: %s", stream, err.Error()))
			return
		}
		if err = consumer.Close(); err != nil {
			log.Debug(ctx, fmt.Sprintf("rabbitmq streams consumer - failed to close consumer from stream %s: %s", stream, err.Error()))
			return
		}
		close(closeChannel)
	}
	log.Debug(ctx, fmt.Sprintf("rabbitmq streams consumer - disconected from stream %s", stream))
}

//

type MessageProcessor func(ctx context.Context, consumerContext stream.ConsumerContext, listener Listener, message *amqp.Message)

func messageProcessor(ctx context.Context, consumerContext stream.ConsumerContext, listener Listener, message *amqp.Message) {
	log.Debug(ctx, fmt.Sprintf("rabbitmq streams consumer - message received: %s", message.Data))
	ctx = context.WithValue(ctx, stream.ConsumerContext{}, consumerContext)
	if err := listener.OnMessage(ctx, message); err != nil {
		log.Debug(ctx, fmt.Sprintf("rabbitmq streams consumer - failed to process message: %s", err.Error()))
	}
}
