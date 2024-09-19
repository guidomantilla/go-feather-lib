package messaging

import (
	"fmt"
)

type BasicMessage[T any] struct {
	headers Headers
	payload T
}

func NewBasicMessage[T any](headers Headers, payload T) *BasicMessage[T] {
	return &BasicMessage[T]{
		headers: headers,
		payload: payload,
	}
}

func (message *BasicMessage[T]) Headers() Headers {
	return message.headers
}

func (message *BasicMessage[T]) Payload() T {
	return message.payload
}

func (message *BasicMessage[T]) String() string {
	return fmt.Sprintf("headers:%v, payload:%T[%v]", message.headers, message.payload, message.payload)
}

//

type BasicErrorPayload struct {
	code    string
	message string
	errors  []string
}

func NewBasicErrorPayload(code string, message string, errors ...string) *BasicErrorPayload {
	return &BasicErrorPayload{
		code:    code,
		message: message,
		errors:  errors,
	}
}

func (payload *BasicErrorPayload) Code() string {
	return payload.code
}

func (payload *BasicErrorPayload) Message() string {
	return payload.message
}

func (payload *BasicErrorPayload) Errors() []string {
	return payload.errors
}

func (payload *BasicErrorPayload) Error() string {
	return fmt.Sprintf("code:%v, message:%v, errors:%v", payload.code, payload.message, payload.errors)
}

func (payload *BasicErrorPayload) String() string {
	return fmt.Sprintf("code:%v, message:%v, errors:%v", payload.code, payload.message, payload.errors)
}

//

type BasicErrorMessage[T any] struct {
	headers Headers
	payload ErrorPayload
	message Message[T]
}

func NewBasicErrorMessage[T any](headers Headers, payload ErrorPayload, message Message[T]) *BasicErrorMessage[T] {
	return &BasicErrorMessage[T]{
		headers: headers,
		payload: payload,
		message: message,
	}
}

func (message *BasicErrorMessage[T]) Headers() Headers {
	return message.headers
}

func (message *BasicErrorMessage[T]) Payload() ErrorPayload {
	return message.payload
}

func (message *BasicErrorMessage[T]) Message() Message[T] {
	return message.message
}

func (message *BasicErrorMessage[T]) String() string {
	return fmt.Sprintf("headers:%v, payload:%T[%v], message:%T[%v]", message.headers, message.payload, message.payload, message.message, message.message)
}
