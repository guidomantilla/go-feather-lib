package integration

import (
	"context"
	"time"

	"github.com/guidomantilla/go-feather-lib/pkg/integration/messaging"
)

func BaseReceiverChannel[T any](name string, handler messaging.ReceiverHandler[T]) messaging.ReceiverChannel[T] {
	return HeadersValidatorReceiverChannel(name, handler)
}

func BaseSenderChannel[T any](name string, handler messaging.SenderHandler[T]) messaging.SenderChannel[T] {
	return HeadersValidatorSenderChannel(name, handler)
}

//

func NullReceiverHandler[T any](ctx context.Context, timeout time.Duration) (messaging.Message[T], error) {
	return nil, nil
}

func NullReceiverChannel[T any](name string) messaging.ReceiverChannel[T] {
	return messaging.NewFunctionAdapterReceiverChannel(name, NullReceiverHandler[T])
}

func LoggedReceiverChannel[T any](name string, handler messaging.ReceiverHandler[T]) messaging.ReceiverChannel[T] {
	return messaging.NewLoggedReceiverChannel(name, messaging.NewFunctionAdapterReceiverChannel(name, handler))
}

func TimeoutReceiverChannel[T any](name string, handler messaging.ReceiverHandler[T]) messaging.ReceiverChannel[T] {
	return messaging.NewTimeoutReceiverChannel(name, LoggedReceiverChannel(name, handler))
}

func HeadersValidatorReceiverChannel[T any](name string, handler messaging.ReceiverHandler[T]) messaging.ReceiverChannel[T] {
	return messaging.NewHeadersValidatorReceiverChannel(name, TimeoutReceiverChannel(name, handler), NullHeadersValidatorValidator())
}

//

func NullSenderHandler[T any](ctx context.Context, message messaging.Message[T], timeout time.Duration) error {
	return nil
}

func NullSenderChannel[T any](name string) messaging.SenderChannel[T] {
	return messaging.NewFunctionAdapterSenderChannel(name, NullSenderHandler[T])
}

func LoggedSenderChannel[T any](name string, handler messaging.SenderHandler[T]) messaging.SenderChannel[T] {
	return messaging.NewLoggedSenderChannel(name, messaging.NewFunctionAdapterSenderChannel(name, handler))
}

func TimeoutSenderChannel[T any](name string, handler messaging.SenderHandler[T]) messaging.SenderChannel[T] {
	return messaging.NewTimeoutSenderChannel(name, LoggedSenderChannel(name, handler))
}

func HeadersValidatorSenderChannel[T any](name string, handler messaging.SenderHandler[T]) messaging.SenderChannel[T] {
	return messaging.NewHeadersValidatorSenderChannel(name, TimeoutSenderChannel(name, handler), NullHeadersValidatorValidator())
}

//

func NullHeadersValidatorHandler(ctx context.Context, headers messaging.Headers) error {
	return nil
}

func NullHeadersValidatorValidator() messaging.HeadersValidator {
	return messaging.NewFunctionAdapterHeadersValidator(NullHeadersValidatorHandler)
}
