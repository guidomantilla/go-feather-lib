package messaging

import (
	"context"
	"time"
)

type BaseSenderChannel[T any] struct {
	handler SenderHandler[T]
}

func NewBaseSenderChannel[T any](handler SenderHandler[T]) *BaseSenderChannel[T] {
	return &BaseSenderChannel[T]{
		handler: handler,
	}
}

func (channel *BaseSenderChannel[T]) Send(ctx context.Context, message Message[T], timeout time.Duration) error {
	return channel.handler(ctx, message, timeout)
}
