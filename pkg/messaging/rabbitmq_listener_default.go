package messaging

import (
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type DefaultRabbitMQListener struct {
}

func NewDefaultRabbitMQListener() *DefaultRabbitMQListener {
	return &DefaultRabbitMQListener{}
}

func (listener *DefaultRabbitMQListener) OnMessage(message *amqp.Delivery) error {

	log.Info(fmt.Sprintf("Received a message: %s", message.Body))
	<-time.After(5 * time.Second)
	return nil
}
