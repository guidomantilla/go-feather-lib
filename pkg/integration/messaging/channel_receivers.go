package messaging

import (
	"context"
	"fmt"
	"time"

	"github.com/guidomantilla/go-feather-lib/pkg/common/assert"
	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type LoggedReceiverChannel[T any] struct {
	name    string
	handler ReceiverHandler[T]
}

func NewLoggedReceiverChannel[T any](name string, handler ReceiverHandler[T]) *LoggedReceiverChannel[T] {
	assert.NotEmpty(name, fmt.Sprintf("integration messaging: %s error - name is required", name))
	assert.NotNil(handler, fmt.Sprintf("integration messaging: %s error - handler is required", name))
	return &LoggedReceiverChannel[T]{
		name:    name,
		handler: handler,
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
	log.Debug(fmt.Sprintf("integration messaging: %s receiving message", channel.name))
	if message, err = channel.handler(ctx, timeout); err != nil {
		log.Debug(fmt.Sprintf("integration messaging: %s error - message not received", channel.name))
		return nil, err
	}

	log.Debug(fmt.Sprintf("integration messaging: %s message received: %v", channel.name, message))
	return message, nil
}

func (channel *LoggedReceiverChannel[T]) Name() string {
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

	errChan := make(chan *response[T])
	go func(errChan chan *response[T]) {
		defer close(errChan)
		errChan <- convert(channel.receiver.Receive(ctx, timeout))
	}(errChan)

	select {
	case <-ctx.Done():
		log.Debug(fmt.Sprintf("integration messaging: %s error - message receiving timeout: %v", channel.name, ctx.Err().Error()))
		return nil, fmt.Errorf("message receiving timeout: %v", ctx.Err().Error())
	case response := <-errChan:
		return response.message, response.err
	}
}

func (channel *TimeoutReceiverChannel[T]) Name() string {
	return channel.name
}

//

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
