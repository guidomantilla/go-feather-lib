package server

import (
	"github.com/qmdx00/lifecycle"
)

var (
	_ Dispatcher = (*NatsMessageDispatcher)(nil)
	_ Dispatcher = (*RabbitMQMessageDispatcher)(nil)
	_ Server     = (*CronServer)(nil)
	_ Server     = (*GrpcServer)(nil)
	_ Server     = (*HttpServer)(nil)
	_ Server     = (*MockServer)(nil)
)

type Dispatcher interface {
	lifecycle.Server
	Dispatch(message any)
}

type Server interface {
	lifecycle.Server
}
