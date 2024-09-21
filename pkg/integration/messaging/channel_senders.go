package messaging

import (
	"context"
	"fmt"
	"time"

	"github.com/guidomantilla/go-feather-lib/pkg/common/assert"
	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type FunctionAdapterSenderChannel[T any] struct {
	name    string
	handler SenderHandler[T]
}

func NewFunctionAdapterSenderChannel[T any](name string, handler SenderHandler[T]) *FunctionAdapterSenderChannel[T] {
	assert.NotEmpty(name, fmt.Sprintf("integration messaging: %s error - name is required", name))
	assert.NotNil(handler, fmt.Sprintf("integration messaging: %s error - handler is required", name))
	return &FunctionAdapterSenderChannel[T]{
		name:    name,
		handler: handler,
	}
}

func (handler *FunctionAdapterSenderChannel[T]) Send(ctx context.Context, timeout time.Duration, message Message[T]) error {
	return handler.handler(ctx, timeout, message)
}

func (handler *FunctionAdapterSenderChannel[T]) Name() string {
	return handler.name
}

//

type HeadersValidatorSenderChannel[T any] struct {
	name       string
	sender     SenderChannel[T]
	validators []HeadersValidator
}

func NewHeadersValidatorSenderChannel[T any](name string, sender SenderChannel[T], validators ...HeadersValidator) *HeadersValidatorSenderChannel[T] {
	assert.NotEmpty(name, fmt.Sprintf("integration messaging: %s error - name is required", name))
	assert.NotNil(sender, fmt.Sprintf("integration messaging: %s error - sender is required", name))
	assert.NotNil(validators, fmt.Sprintf("integration messaging: %s error - validators are required", name))
	assert.NotEmpty(validators, fmt.Sprintf("integration messaging: %s error - validators are required", name))
	return &HeadersValidatorSenderChannel[T]{
		name:       name,
		sender:     sender,
		validators: validators,
	}
}

func (channel *HeadersValidatorSenderChannel[T]) Send(ctx context.Context, timeout time.Duration, message Message[T]) error {

	if ctx == nil {
		return fmt.Errorf("integration messaging: %s error - for sending a message, context is required", channel.name)
	}

	if message == nil {
		return fmt.Errorf("integration messaging: %s error - for sending a message, message is required", channel.name)
	}

	var err error
	for _, validator := range channel.validators {
		if err = validator.Validate(ctx, message.Headers()); err != nil {
			return err
		}
	}

	if err = channel.sender.Send(ctx, timeout, message); err != nil {
		return err
	}

	return nil
}

func (channel *HeadersValidatorSenderChannel[T]) Name() string {
	return channel.name
}

//

type TimeoutSenderChannel[T any] struct {
	name   string
	sender SenderChannel[T]
}

func NewTimeoutSenderChannel[T any](name string, sender SenderChannel[T]) *TimeoutSenderChannel[T] {
	assert.NotEmpty(name, fmt.Sprintf("integration messaging: %s error - name is required", name))
	assert.NotNil(sender, fmt.Sprintf("integration messaging: %s error - sender is required", name))
	return &TimeoutSenderChannel[T]{
		name:   name,
		sender: sender,
	}
}

func (channel *TimeoutSenderChannel[T]) Send(ctx context.Context, timeout time.Duration, message Message[T]) error {

	if ctx == nil {
		return fmt.Errorf("integration messaging: %s error - for sending a message, context is required", channel.name)
	}

	if message == nil {
		return fmt.Errorf("integration messaging: %s error - for sending a message, message is required", channel.name)
	}

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	errChan := make(chan error)
	go func(errChan chan error) {
		defer close(errChan)
		errChan <- channel.sender.Send(ctx, timeout, message)
	}(errChan)

	select {
	case <-ctx.Done():
		message.Headers().Add(HeaderExpired, "true")
		message.Headers().Add("x-error-detail", ctx.Err().Error())
		return fmt.Errorf("message sending timeout: %v", ctx.Err().Error())
	case err := <-errChan:
		return err
	}
}

func (channel *TimeoutSenderChannel[T]) Name() string {
	return channel.name
}

//

type LoggedSenderChannel[T any] struct {
	name   string
	sender SenderChannel[T]
}

func NewLoggedSenderChannel[T any](name string, sender SenderChannel[T]) *LoggedSenderChannel[T] {
	assert.NotEmpty(name, fmt.Sprintf("integration messaging: %s error - name is required", name))
	assert.NotNil(sender, fmt.Sprintf("integration messaging: %s error - sender is required", name))
	return &LoggedSenderChannel[T]{
		name:   name,
		sender: sender,
	}
}

func (channel *LoggedSenderChannel[T]) Send(ctx context.Context, timeout time.Duration, message Message[T]) error {

	if ctx == nil {
		return fmt.Errorf("integration messaging: %s error - for sending a message, context is required", channel.name)
	}

	if message == nil {
		return fmt.Errorf("integration messaging: %s error - for sending a message, message is required", channel.name)
	}

	log.Trace(fmt.Sprintf("integration messaging: sending message: %v", message))
	if err := channel.sender.Send(ctx, timeout, message); err != nil {
		log.Trace(fmt.Sprintf("integration messaging: error - message not sent: %v", message))
		return err
	}

	log.Trace(fmt.Sprintf("integration messaging: message sent: %v", message))
	return nil
}

func (channel *LoggedSenderChannel[T]) Name() string {
	return channel.name
}
