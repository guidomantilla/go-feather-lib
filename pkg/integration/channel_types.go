package integration

import (
	"context"
	"time"

	"github.com/guidomantilla/go-feather-lib/pkg/integration/messaging"
)

func BaseReceiverChannel[T any](name string, handler messaging.ReceiverHandler[T]) messaging.ReceiverChannel[T] {
	var channel messaging.ReceiverChannel[T]
	channel = messaging.NewFunctionAdapterReceiverChannel(name, handler)
	channel = messaging.NewHeadersValidatorReceiverChannel(name, channel, NullHeadersValidatorValidator())
	channel = messaging.NewTimeoutReceiverChannel(name, channel)
	channel = messaging.NewLoggedReceiverChannel(name, channel)
	return channel
}

func BaseSenderChannel[T any](name string, handler messaging.SenderHandler[T]) messaging.SenderChannel[T] {
	var channel messaging.SenderChannel[T]
	channel = messaging.NewFunctionAdapterSenderChannel(name, handler)
	channel = messaging.NewHeadersValidatorSenderChannel(name, channel, NullHeadersValidatorValidator())
	channel = messaging.NewTimeoutSenderChannel(name, channel)
	channel = messaging.NewLoggedSenderChannel(name, channel)
	return channel
}

//

func NullReceiverHandler[T any](ctx context.Context, timeout time.Duration) (messaging.Message[T], error) {
	return nil, nil
}

func NullReceiverChannel[T any](name string) messaging.ReceiverChannel[T] {
	return messaging.NewFunctionAdapterReceiverChannel(name, NullReceiverHandler[T])
}

//

func NullSenderHandler[T any](ctx context.Context, timeout time.Duration, message messaging.Message[T]) error {
	return nil
}

func NullSenderChannel[T any](name string) messaging.SenderChannel[T] {
	return messaging.NewFunctionAdapterSenderChannel(name, NullSenderHandler[T])
}

//

func NullHeadersValidatorHandler(ctx context.Context, headers messaging.Headers) error {
	return nil
}

func NullHeadersValidatorValidator() messaging.HeadersValidator {
	return messaging.NewFunctionAdapterHeadersValidator(NullHeadersValidatorHandler)
}
