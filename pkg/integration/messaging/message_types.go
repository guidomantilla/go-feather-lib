package messaging

import (
	"context"
	"fmt"
	"time"
)

type Message[T any] interface {
	fmt.Stringer
	Headers() Headers
	Payload() T
}

type ErrorPayload interface {
	fmt.Stringer
	error
	Code() string
	Message() string
	Errors() []string
}

type ErrorMessage[T any] interface {
	fmt.Stringer
	Headers() Headers
	Payload() ErrorPayload
	Message() Message[T]
}

//

type MessageStream[T any] chan Message[T]

type MessagePipeline[T any] struct {
	name      string
	internal  chan Message[T]
	closeChan chan struct{}
}

func NewMessagePipeline[T any](name string, size int) *MessagePipeline[T] {
	return &MessagePipeline[T]{
		name:      name,
		internal:  make(chan Message[T], size),
		closeChan: make(chan struct{}),
	}
}

func (pipeline *MessagePipeline[T]) Send(ctx context.Context, timeout time.Duration, message Message[T]) error {

	if ctx == nil {
		return fmt.Errorf("integration messaging: %s error - for sending a message, context is required", pipeline.name)
	}

	if message == nil {
		return fmt.Errorf("integration messaging: %s error - for sending a message, message is required", pipeline.name)
	}

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	select {
	case <-ctx.Done():
		message.Headers().Add(HeaderExpired, "true")
		message.Headers().Add("x-error-detail", ctx.Err().Error())
		return fmt.Errorf("message pipeline timeout: %v", ctx.Err().Error())
	case <-pipeline.closeChan:
		return fmt.Errorf("message pipeline %s is closed", pipeline.name)
	default:
		pipeline.internal <- message
		return nil
	}
}

func (pipeline *MessagePipeline[T]) Receive(ctx context.Context, timeout time.Duration) (Message[T], error) {

	if ctx == nil {
		return nil, fmt.Errorf("integration messaging: %s error - for receiving a message, context is required", pipeline.name)
	}

	if timeout <= 0 {
		return nil, fmt.Errorf("integration messaging: %s error - for receiving a message, timeout is required", pipeline.name)
	}

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("message pipeline timeout: %v", ctx.Err().Error())
	case <-pipeline.closeChan:
		return nil, fmt.Errorf("message pipeline %s is closed", pipeline.name)
	case message := <-pipeline.internal:
		return message, nil
	}
}

func (pipeline *MessagePipeline[T]) Close() {
	pipeline.closeChan <- struct{}{}
	time.Sleep(2 * time.Second)
	close(pipeline.internal)
	close(pipeline.closeChan)
}

func (pipeline *MessagePipeline[T]) Name() string {
	return pipeline.name
}
