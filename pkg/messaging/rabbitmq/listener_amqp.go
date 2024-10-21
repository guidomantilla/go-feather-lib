package rabbitmq

import (
	"context"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type AmqpListener struct {
}

func NewAmqpListener() *AmqpListener {
	return &AmqpListener{}
}

func (listener *AmqpListener) OnMessage(ctx context.Context, message *amqp.Delivery) error {

	log.Info(fmt.Sprintf("Received a message: %s", message.Body))
	time.Sleep(5 * time.Second)
	return nil
}
