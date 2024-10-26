package server

import (
	"net/http"

	cron "github.com/robfig/cron/v3"
	"google.golang.org/grpc"

	"github.com/guidomantilla/go-feather-lib/pkg/messaging/rabbitmq"
)

func BuildBaseServer() (string, Server) {
	return "base-server", NewBaseServer()
}

func BuildCronServer(cron *cron.Cron) (string, Server) {
	return "cron-server", NewCronServer(cron)
}

func BuildHttpServer(server *http.Server) (string, Server) {
	return "http-server", NewHttpServer(server)
}

func BuildGrpcServer(address string, server *grpc.Server) (string, Server) {
	return "grpc-server", NewGrpcServer(address, server)
}

func BuildRabbitMQServer(consumers ...rabbitmq.Consumer) (string, Server) {
	return "rabbitmq-server", NewRabbitMQServer(consumers...)
}
