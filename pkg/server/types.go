package server

import (
	"github.com/qmdx00/lifecycle"
)

var (
	_ Server = (*RabbitMQServer)(nil)
	_ Server = (*CronServer)(nil)
	_ Server = (*GrpcServer)(nil)
	_ Server = (*HttpServer)(nil)
	_ Server = (*MockServer)(nil)
)

type Server interface {
	lifecycle.Server
}
