package messaging

import (
	"context"
	"fmt"
	"time"

	"github.com/guidomantilla/go-feather-lib/pkg/common/assert"
)

// PassThroughChannel

type PassThroughChannel[T any] struct {
	name     string
	sender   SenderChannel[T]
	receiver ReceiverChannel[T]
}

func NewPassThroughChannel[T any](name string, sender SenderChannel[T], receiver ReceiverChannel[T]) *PassThroughChannel[T] {
	assert.NotEmpty(name, fmt.Sprintf("integration messaging: %s error - name is required", name))
	assert.NotNil(sender, fmt.Sprintf("integration messaging: %s error - sender is required", name))
	assert.NotNil(receiver, fmt.Sprintf("integration messaging: %s error - receiver is required", name))
	return &PassThroughChannel[T]{
		name:     name,
		sender:   sender,
		receiver: receiver,
	}
}

func (channel *PassThroughChannel[T]) Send(ctx context.Context, timeout time.Duration, message Message[T]) error {

	if ctx == nil {
		return fmt.Errorf("integration messaging: %s error - for sending a message, context is required", channel.name)
	}

	if message == nil {
		return fmt.Errorf("integration messaging: %s error - for sending a message, message is required", channel.name)
	}

	return channel.sender.Send(ctx, timeout, message)
}

func (channel *PassThroughChannel[T]) Receive(ctx context.Context, timeout time.Duration) (Message[T], error) {

	if ctx == nil {
		return nil, fmt.Errorf("integration messaging: %s error - for receiving a message, context is required", channel.name)
	}

	if timeout <= 0 {
		return nil, fmt.Errorf("integration messaging: %s error - for receiving a message, timeout is required", channel.name)
	}

	return channel.receiver.Receive(ctx, timeout)
}

func (channel *PassThroughChannel[T]) Name() string {
	return channel.name
}

// QueueChannel

type QueueChannel[T any] struct {
	name      string
	internal  chan Message[T]
	closeChan chan struct{}
}

func NewQueueChannel[T any](name string, size int) *QueueChannel[T] {
	return &QueueChannel[T]{
		name:      name,
		internal:  make(chan Message[T], size),
		closeChan: make(chan struct{}),
	}
}

func (channel *QueueChannel[T]) Send(ctx context.Context, timeout time.Duration, message Message[T]) error {

	if ctx == nil {
		return fmt.Errorf("integration messaging: %s error - for sending a message, context is required", channel.name)
	}

	if message == nil {
		return fmt.Errorf("integration messaging: %s error - for sending a message, message is required", channel.name)
	}

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	select {
	case <-ctx.Done():
		message.Headers().Add(HeaderExpired, "true")
		message.Headers().Add("x-error-detail", ctx.Err().Error())
		return fmt.Errorf("message pipeline timeout: %v", ctx.Err().Error())
	case <-channel.closeChan:
		return fmt.Errorf("message pipeline %s is closed", channel.name)
	default:
		channel.internal <- message
		return nil
	}
}

func (channel *QueueChannel[T]) Receive(ctx context.Context, timeout time.Duration) (Message[T], error) {

	if ctx == nil {
		return nil, fmt.Errorf("integration messaging: %s error - for receiving a message, context is required", channel.name)
	}

	if timeout <= 0 {
		return nil, fmt.Errorf("integration messaging: %s error - for receiving a message, timeout is required", channel.name)
	}

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("message pipeline timeout: %v", ctx.Err().Error())
	case <-channel.closeChan:
		return nil, fmt.Errorf("message pipeline %s is closed", channel.name)
	case message := <-channel.internal:
		return message, nil
	}
}

func (channel *QueueChannel[T]) Close() {
	channel.closeChan <- struct{}{}
	time.Sleep(2 * time.Second)
	close(channel.internal)
	close(channel.closeChan)
}

func (channel *QueueChannel[T]) Name() string {
	return channel.name
}
