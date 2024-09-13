package messaging

import (
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type DefaultRabbitMQMessageListener struct {
}

func NewDefaultRabbitMQMessageListener() *DefaultRabbitMQMessageListener {
	return &DefaultRabbitMQMessageListener{}
}

func (listener *DefaultRabbitMQMessageListener) OnMessage(message *amqp.Delivery) error {
	log.Info(fmt.Sprintf("Received a message: %s", message.Body))
	<-time.After(5 * time.Second)
	return nil
}
