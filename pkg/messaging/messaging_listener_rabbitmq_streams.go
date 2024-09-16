package messaging

import (
	"context"
	"fmt"
	"time"

	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/amqp"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type RabbitMQStreamsListener struct {
}

func NewRabbitMQStreamsListener() *RabbitMQStreamsListener {
	return &RabbitMQStreamsListener{}
}

func (listener *RabbitMQStreamsListener) OnMessage(ctx context.Context, message *amqp.Message) error {

	log.Info(fmt.Sprintf("Received a message: %s", message.Data))
	<-time.After(5 * time.Second)
	return nil
}
