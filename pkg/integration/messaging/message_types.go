package messaging

import (
	"fmt"
)

type MessageStream[T any] chan Message[T]

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
