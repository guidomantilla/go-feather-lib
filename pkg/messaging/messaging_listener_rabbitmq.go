package messaging

import (
	"context"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type RabbitMQListener struct {
}

func NewRabbitMQListener() *RabbitMQListener {
	return &RabbitMQListener{}
}

func (listener *RabbitMQListener) OnMessage(ctx context.Context, message *amqp.Delivery) error {

	log.Info(fmt.Sprintf("Received a message: %s", message.Body))
	<-time.After(5 * time.Second)
	return nil
}
