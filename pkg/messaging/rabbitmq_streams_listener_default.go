package messaging

import (
	"fmt"
	"time"

	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/amqp"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type DefaultRabbitMQStreamsListener struct {
}

func NewDefaultRabbitMQStreamsListener() *DefaultRabbitMQStreamsListener {
	return &DefaultRabbitMQStreamsListener{}
}

func (listener *DefaultRabbitMQStreamsListener) OnMessage(message *amqp.Message) error {

	log.Info(fmt.Sprintf("Received a message: %s", message.Data))
	<-time.After(5 * time.Second)
	return nil
}
