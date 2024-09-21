package messaging

import (
	"context"
	"time"
)

type SenderHandler[T any] func(ctx context.Context, timeout time.Duration, message Message[T]) error

type SenderChannel[T any] interface {
	Send(ctx context.Context, timeout time.Duration, message Message[T]) error
	Name() string
}

//

type ReceiverHandler[T any] func(ctx context.Context, timeout time.Duration) (Message[T], error)

type ReceiverChannel[T any] interface {
	Receive(ctx context.Context, timeout time.Duration) (Message[T], error)
	Name() string
}

//

type MessageHandler[T any] func(ctx context.Context, message Message[T]) error

type MessageChannel[T any] interface {
	SenderChannel[T]
	ReceiverChannel[T]
}

//

type ProducerHandler[T any] func(ctx context.Context, timeout time.Duration, stream MessagePipeline[T]) error

type ConsumerHandler[T any] func(ctx context.Context, timeout time.Duration) (MessagePipeline[T], error)

//
