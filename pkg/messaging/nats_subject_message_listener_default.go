package messaging

import (
	"fmt"
	"time"

	nats "github.com/nats-io/nats.go"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type DefaultNatsMessageListener struct {
}

func NewDefaultNatsMessageListener() *DefaultNatsMessageListener {
	return &DefaultNatsMessageListener{}
}

func (listener *DefaultNatsMessageListener) OnMessage(message *nats.Msg) error {
	log.Info(fmt.Sprintf("Received a message: %s", message.Data))
	<-time.After(5 * time.Second)
	return nil
}
