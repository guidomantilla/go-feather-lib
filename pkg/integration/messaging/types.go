package messaging

import (
	"context"
	"time"
)

//

type MoreComplexSenderChannel[T any] struct {
	SenderChannel[T]
}

func NewMoreComplexSenderChannel[T any](handler SenderHandler[T]) *MoreComplexSenderChannel[T] {
	return &MoreComplexSenderChannel[T]{
		SenderChannel: NewBaseSenderChannel[T](handler),
	}
}

func (channel *MoreComplexSenderChannel[T]) Send(ctx context.Context, message Message[T], timeout time.Duration) error {
	return channel.SenderChannel.Send(ctx, message, timeout)
}
