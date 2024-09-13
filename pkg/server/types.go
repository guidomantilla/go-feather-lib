package server

import (
	"github.com/qmdx00/lifecycle"
)

var (
	_ Dispatcher = (*NatsMessageDispatcher)(nil)
	_ Dispatcher = (*RabbitMQQueueMessageDispatcher)(nil)
	_ Server     = (*CronServer)(nil)
	_ Server     = (*GrpcServer)(nil)
	_ Server     = (*HttpServer)(nil)
)

type Dispatcher interface {
	lifecycle.Server
	ListenAndDispatch() error
	Dispatch(message any)
}

type Server interface {
	lifecycle.Server
}
