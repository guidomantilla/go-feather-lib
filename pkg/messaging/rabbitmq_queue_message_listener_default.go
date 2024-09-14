package messaging

import (
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type DefaultRabbitMQQueueMessageListener struct {
}

func NewDefaultRabbitMQQueueMessageListener() *DefaultRabbitMQQueueMessageListener {
	return &DefaultRabbitMQQueueMessageListener{}
}

func (listener *DefaultRabbitMQQueueMessageListener) OnMessage(message *amqp.Delivery) error {

	if message.Body == nil || len(message.Body) == 0 || string(message.Body) == "" {
		return nil
	}

	log.Info(fmt.Sprintf("Received a message: %s", message.Body))
	<-time.After(5 * time.Second)
	return nil
}
