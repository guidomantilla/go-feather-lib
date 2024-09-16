package main

import (
	"os"
	"syscall"

	"github.com/qmdx00/lifecycle"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
	"github.com/guidomantilla/go-feather-lib/pkg/messaging"
	"github.com/guidomantilla/go-feather-lib/pkg/server"
)

func main() {

	var err error
	appName, version := "rabbitmq-stream-micro-stream", "1.0.0"
	os.Setenv("LOG_LEVEL", "DEBUG")
	log.Custom()

	app := lifecycle.NewApp(
		lifecycle.WithName(appName), lifecycle.WithVersion(version),
		lifecycle.WithSignal(syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGKILL),
	)

	messagingContext := messaging.NewDefaultMessagingContext("rabbitmq-stream://:username::password@:server:vhost",
		"raven-dev", "raven-dev*+", "170.187.157.212:5552", messaging.WithVhost("/"))
	connection := messaging.NewRabbitMQConnection(messagingContext, messaging.WithRabbitMQStreamsDialer())

	streams := messaging.NewRabbitMQStreams(connection, "rabbitmq-stream-micro-stream")
	app.Attach("RabbitMQServer", server.BuildRabbitMQServer(streams))

	if err = app.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
