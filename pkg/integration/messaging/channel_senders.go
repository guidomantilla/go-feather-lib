package messaging

import (
	"context"
	"fmt"
	"time"

	"github.com/guidomantilla/go-feather-lib/pkg/common/assert"
	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type LoggedSenderChannel[T any] struct {
	name    string
	handler SenderHandler[T]
}

func NewLoggedSenderChannel[T any](name string, handler SenderHandler[T]) *LoggedSenderChannel[T] {
	assert.NotEmpty(name, fmt.Sprintf("integration messaging: %s error - name is required", name))
	assert.NotNil(handler, fmt.Sprintf("integration messaging: %s error - handler is required", name))
	return &LoggedSenderChannel[T]{
		name:    name,
		handler: handler,
	}
}

func (channel *LoggedSenderChannel[T]) Send(ctx context.Context, message Message[T], timeout time.Duration) error {

	if ctx == nil {
		return fmt.Errorf("integration messaging: %s error - for sending a message, context is required", channel.name)
	}

	if message == nil {
		return fmt.Errorf("integration messaging: %s error - for sending a message, message is required", channel.name)
	}

	log.Debug(fmt.Sprintf("integration messaging: sending message: %v", message))
	if err := channel.handler(ctx, message, timeout); err != nil {
		log.Debug(fmt.Sprintf("integration messaging: error - message not sent: %v", message))
		return err
	}

	log.Debug(fmt.Sprintf("integration messaging: message sent: %v", message))
	return nil
}

func (channel *LoggedSenderChannel[T]) Name() string {
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

func (channel *TimeoutSenderChannel[T]) Send(ctx context.Context, message Message[T], timeout time.Duration) error {

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
		errChan <- channel.sender.Send(ctx, message, timeout)
	}(errChan)

	select {
	case <-ctx.Done():
		log.Debug(fmt.Sprintf("integration messaging: %s error - message sending timeout: %v", channel.name, ctx.Err().Error()))
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
