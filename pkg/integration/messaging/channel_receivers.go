package messaging

import (
	"context"
	"time"
)

type BaseReceiverChannel[T any] struct {
	handler ReceiverHandler[T]
}

func NewBaseReceiverChannel[T any](handler ReceiverHandler[T]) *BaseReceiverChannel[T] {
	return &BaseReceiverChannel[T]{
		handler: handler,
	}
}

func (channel *BaseReceiverChannel[T]) Receive(ctx context.Context, timeout time.Duration) (Message[T], error) {
	return channel.handler(ctx, timeout)
}

func (channel *BaseReceiverChannel[T]) Name() string {
	return "base-receiver-channel"
}
