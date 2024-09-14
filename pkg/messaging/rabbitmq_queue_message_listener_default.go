package messaging

import (
	"fmt"
	"reflect"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

var JunkMessage = amqp.Delivery{}

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

	if reflect.DeepEqual(*message, JunkMessage) {
		return nil
	}

	log.Info(fmt.Sprintf("Received a message: %s", message.Body))
	<-time.After(5 * time.Second)
	return nil
}
