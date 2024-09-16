package messaging

import (
	"context"
	"fmt"
	"time"

	nats "github.com/nats-io/nats.go"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type NatsListener struct {
}

func NewNatsListener() *NatsListener {
	return &NatsListener{}
}

func (listener *NatsListener) OnMessage(ctx context.Context, message *nats.Msg) error {
	log.Info(fmt.Sprintf("Received a message: %s", message.Data))
	<-time.After(5 * time.Second)
	return nil
}
