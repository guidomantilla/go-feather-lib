package amqp

import (
	"context"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type listener struct {
}

func NewListener() *listener {
	return &listener{}
}

func (listener *listener) OnMessage(ctx context.Context, message *amqp.Delivery) error {

	log.Info(ctx, fmt.Sprintf("Received a message: %s", message.Body))
	time.Sleep(5 * time.Second)
	return nil
}
