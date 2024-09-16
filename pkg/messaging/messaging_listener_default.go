package messaging

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type DefaultMessagingListener[T MessagingListenerTypes] struct {
}

func NewDefaultMessagingListener[T MessagingListenerTypes]() *DefaultMessagingListener[T] {
	return &DefaultMessagingListener[T]{}
}

func (listener *DefaultMessagingListener[T]) OnMessage(ctx context.Context, message T) error {

	var err error

	var jsonBytes []byte
	if jsonBytes, err = json.Marshal(message); err != nil {
		return err
	}

	log.Info(fmt.Sprintf("Received a message: %s", string(jsonBytes)))
	<-time.After(5 * time.Second)
	return nil
}
