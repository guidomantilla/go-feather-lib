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
	appName, version := "rabbitmq-micro", "1.0.0"
	os.Setenv("LOG_LEVEL", "DEBUG")
	log.Custom()
	app := lifecycle.NewApp(
		lifecycle.WithName(appName), lifecycle.WithVersion(version),
		lifecycle.WithSignal(syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGKILL),
	)

	messagingContext := messaging.NewDefaultRabbitMQContext("amqp://:username::password@:server/", "raven-dev", "raven-dev*+", "170.187.157.212:5672")
	connection := messaging.NewDefaultRabbitMQConnection(messagingContext)
	defer connection.Close()

	connection.Connect()
	app.Attach("DummyServer", server.BuildDummyServer())

	//listener := messaging.NewDefaultRabbitMQQueueMessageListener("my-queue")
	//dispatcher := server.BuildRabbitMQQueueMessageDispatcher(messagingContext, connection, listener)
	//app.Attach("RabbitMQDispatcher", dispatcher)

	if err = app.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
