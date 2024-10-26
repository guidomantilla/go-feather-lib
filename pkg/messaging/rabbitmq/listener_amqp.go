package rabbitmq

import (
	"context"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type amqpListener struct {
}

func NewAmqpListener() *amqpListener {
	return &amqpListener{}
}

func (listener *amqpListener) OnMessage(ctx context.Context, message *amqp.Delivery) error {

	log.Info(fmt.Sprintf("Received a message: %s", message.Body))
	time.Sleep(5 * time.Second)
	return nil
}
