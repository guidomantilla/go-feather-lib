package messaging

import (
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type DefaultRabbitMQQueueMessageListener struct {
	queue string
}

func NewDefaultRabbitMQQueueMessageListener(queue string) *DefaultRabbitMQQueueMessageListener {
	return &DefaultRabbitMQQueueMessageListener{
		queue: queue,
	}
}

func (listener *DefaultRabbitMQQueueMessageListener) Queue() string {
	return listener.queue
}

func (listener *DefaultRabbitMQQueueMessageListener) OnMessage(message *amqp.Delivery) error {

	log.Info(fmt.Sprintf("Received a message: %s", message.Body))
	<-time.After(5 * time.Second)
	return nil
}
