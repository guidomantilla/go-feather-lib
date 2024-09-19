package messaging

import (
	"fmt"
)

type BaseMessage[T any] struct {
	headers Headers
	payload T
}

func NewBaseMessage[T any](headers Headers, payload T) *BaseMessage[T] {
	return &BaseMessage[T]{
		headers: headers,
		payload: payload,
	}
}

func (message *BaseMessage[T]) Headers() Headers {
	return message.headers
}

func (message *BaseMessage[T]) Payload() T {
	return message.payload
}

func (message *BaseMessage[T]) String() string {
	return fmt.Sprintf("headers:%v, payload:%T[%v]", message.headers, message.payload, message.payload)
}

//

type BaseErrorPayload struct {
	code    string
	message string
	errors  []string
}

func NewBaseErrorPayload(code string, message string, errors ...string) *BaseErrorPayload {
	return &BaseErrorPayload{
		code:    code,
		message: message,
		errors:  errors,
	}
}

func (payload *BaseErrorPayload) Code() string {
	return payload.code
}

func (payload *BaseErrorPayload) Message() string {
	return payload.message
}

func (payload *BaseErrorPayload) Errors() []string {
	return payload.errors
}

func (payload *BaseErrorPayload) String() string {
	return fmt.Sprintf("code:%v, message:%v, errors:%v", payload.code, payload.message, payload.errors)
}

func (payload *BaseErrorPayload) Error() string {
	return payload.String()
}

//

type BaseErrorMessage[T any] struct {
	headers Headers
	payload ErrorPayload
	message Message[T]
}

func NewBaseErrorMessage[T any](headers Headers, payload ErrorPayload, message Message[T]) *BaseErrorMessage[T] {
	return &BaseErrorMessage[T]{
		headers: headers,
		payload: payload,
		message: message,
	}
}

func (message *BaseErrorMessage[T]) Headers() Headers {
	return message.headers
}

func (message *BaseErrorMessage[T]) Payload() ErrorPayload {
	return message.payload
}

func (message *BaseErrorMessage[T]) Message() Message[T] {
	return message.message
}

func (message *BaseErrorMessage[T]) String() string {
	return fmt.Sprintf("headers:%v, payload:%T[%v], message:%T[%v]", message.headers, message.payload, message.payload, message.message, message.message)
}
