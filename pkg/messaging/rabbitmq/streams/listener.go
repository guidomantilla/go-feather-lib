package streams

import (
	"context"
	"fmt"
	"time"

	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/amqp"
	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/stream"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type listener struct {
}

func NewListener() *listener {
	return &listener{}
}

func (listener *listener) OnMessage(ctx context.Context, message *amqp.Message) error {

	var consumerContext stream.ConsumerContext
	anyContext := ctx.Value(stream.ConsumerContext{})
	if anyContext != nil {
		consumerContext = anyContext.(stream.ConsumerContext)
	}
	log.Debug(fmt.Sprintf("Received a consumerContext: %s", consumerContext.Consumer.GetName()))
	log.Info(fmt.Sprintf("Received a message: %s", message.Data))
	time.Sleep(5 * time.Second)
	return nil
}
