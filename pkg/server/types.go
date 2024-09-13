package server

import (
	"github.com/qmdx00/lifecycle"
)

var _ Dispatcher = (*NatsMessageDispatcher)(nil)

var _ Dispatcher = (*RabbitMQQueueMessageDispatcher)(nil)

type Dispatcher interface {
	lifecycle.Server
	ListenAndDispatch() error
	Dispatch(message any)
}
