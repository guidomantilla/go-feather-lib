package messaging

import (
	"context"
	"fmt"
	"time"

	"github.com/guidomantilla/go-feather-lib/pkg/common/assert"
	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

//

type FunctionAdapterReceiverChannel[T any] struct {
	name    string
	handler ReceiverHandler[T]
}

func NewFunctionAdapterReceiverChannel[T any](name string, handler ReceiverHandler[T]) *FunctionAdapterReceiverChannel[T] {
	assert.NotEmpty(name, fmt.Sprintf("integration messaging: %s error - name is required", name))
	assert.NotNil(handler, fmt.Sprintf("integration messaging: %s error - handler is required", name))
	return &FunctionAdapterReceiverChannel[T]{
		name:    name,
		handler: handler,
	}
}

func (handler *FunctionAdapterReceiverChannel[T]) Receive(ctx context.Context, timeout time.Duration) (Message[T], error) {
	return handler.handler(ctx, timeout)
}

func (handler *FunctionAdapterReceiverChannel[T]) Name() string {
	return handler.name
}

//

type HeadersValidatorReceiverChannel[T any] struct {
	name       string
	receiver   ReceiverChannel[T]
	validators []HeadersValidator
}

func NewHeadersValidatorReceiverChannel[T any](name string, receiver ReceiverChannel[T], validators ...HeadersValidator) *HeadersValidatorReceiverChannel[T] {
	assert.NotEmpty(name, fmt.Sprintf("integration messaging: %s error - name is required", name))
	assert.NotNil(receiver, fmt.Sprintf("integration messaging: %s error - receiver is required", name))
	assert.NotNil(validators, fmt.Sprintf("integration messaging: %s error - validators are required", name))
	assert.NotEmpty(validators, fmt.Sprintf("integration messaging: %s error - validators are required", name))
	return &HeadersValidatorReceiverChannel[T]{
		name:       name,
		receiver:   receiver,
		validators: validators,
	}
}

func (channel *HeadersValidatorReceiverChannel[T]) Receive(ctx context.Context, timeout time.Duration) (Message[T], error) {

	var err error
	var message Message[T]
	if message, err = channel.receiver.Receive(ctx, timeout); err != nil {
		return nil, err
	}

	for _, validator := range channel.validators {
		if err = validator.Validate(ctx, message.Headers()); err != nil {
			return nil, err
		}
	}

	return message, nil

}

func (channel *HeadersValidatorReceiverChannel[T]) Name() string {
	return channel.name
}

//

type TimeoutReceiverChannel[T any] struct {
	name     string
	receiver ReceiverChannel[T]
}

func NewTimeoutReceiverChannel[T any](name string, receiver ReceiverChannel[T]) *TimeoutReceiverChannel[T] {
	assert.NotEmpty(name, fmt.Sprintf("integration messaging: %s error - name is required", name))
	assert.NotNil(receiver, fmt.Sprintf("integration messaging: %s error - receiver is required", name))
	return &TimeoutReceiverChannel[T]{
		name:     name,
		receiver: receiver,
	}
}

func (channel *TimeoutReceiverChannel[T]) Receive(ctx context.Context, timeout time.Duration) (Message[T], error) {

	if ctx == nil {
		return nil, fmt.Errorf("integration messaging: %s error - for receiving a message, context is required", channel.name)
	}

	if timeout <= 0 {
		return nil, fmt.Errorf("integration messaging: %s error - for receiving a message, timeout is required", channel.name)
	}

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	responseChan := make(chan *response[T])
	go func(responseChan chan *response[T]) {
		defer close(responseChan)
		responseChan <- convert(channel.receiver.Receive(ctx, timeout))
	}(responseChan)

	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("message receiving timeout: %v", ctx.Err().Error())
	case response := <-responseChan:
		return response.message, response.err
	}
}

func (channel *TimeoutReceiverChannel[T]) Name() string {
	return channel.name
}

type response[T any] struct {
	message Message[T]
	err     error
}

func convert[T any](message Message[T], err error) *response[T] {
	return &response[T]{
		message: message,
		err:     err,
	}
}

//

type LoggedReceiverChannel[T any] struct {
	name     string
	receiver ReceiverChannel[T]
}

func NewLoggedReceiverChannel[T any](name string, receiver ReceiverChannel[T]) *LoggedReceiverChannel[T] {
	assert.NotEmpty(name, fmt.Sprintf("integration messaging: %s error - name is required", name))
	assert.NotNil(receiver, fmt.Sprintf("integration messaging: %s error - receiver is required", name))
	return &LoggedReceiverChannel[T]{
		name:     name,
		receiver: receiver,
	}
}

func (channel *LoggedReceiverChannel[T]) Receive(ctx context.Context, timeout time.Duration) (Message[T], error) {

	if ctx == nil {
		return nil, fmt.Errorf("integration messaging: %s error - for receiving a message, context is required", channel.name)
	}

	if timeout <= 0 {
		return nil, fmt.Errorf("integration messaging: %s error - for receiving a message, timeout is required", channel.name)
	}

	var err error
	var message Message[T]
	log.Trace(fmt.Sprintf("integration messaging: %s receiving message", channel.name))
	if message, err = channel.receiver.Receive(ctx, timeout); err != nil {
		log.Trace(fmt.Sprintf("integration messaging: %s error - message not received", channel.name))
		return nil, err
	}

	log.Trace(fmt.Sprintf("integration messaging: %s message received: %v", channel.name, message))
	return message, nil
}

func (channel *LoggedReceiverChannel[T]) Name() string {
	return channel.name
}
