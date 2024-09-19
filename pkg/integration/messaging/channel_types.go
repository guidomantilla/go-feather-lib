package messaging

import (
	"context"
	"time"
)

type SenderHandler[T any] func(ctx context.Context, message Message[T], timeout time.Duration) error

type SenderChannel[T any] interface {
	Send(ctx context.Context, message Message[T], timeout time.Duration) error
}

//

type ReceiverHandler[T any] func(ctx context.Context, timeout time.Duration) (Message[T], error)

type ReceiverChannel[T any] interface {
	Receive(ctx context.Context, timeout time.Duration) (Message[T], error)
}

//

type SenderChannelInterceptor[T any] interface {
	AroundSender(ctx context.Context, message Message[T], timeout time.Duration, sender SenderChannel[T]) error
}

//
