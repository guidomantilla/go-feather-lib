package main

import (
	"syscall"

	"github.com/qmdx00/lifecycle"
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
	"github.com/guidomantilla/go-feather-lib/pkg/messaging"
	"github.com/guidomantilla/go-feather-lib/pkg/server"
)

func main() {

	var err error
	appName, version := "rabbitmq-micro", "1.0.0"
	log.Default()
	app := lifecycle.NewApp(
		lifecycle.WithName(appName), lifecycle.WithVersion(version),
		lifecycle.WithSignal(syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGKILL),
	)

	messagingContext := messaging.NewDefaultMessagingContext("amqp://:username::password@:server:vhost",
		"raven-dev", "raven-dev*+", "170.187.157.212:5672", messaging.WithVhost("/"))
	connection := messaging.NewRabbitMQConnection(messagingContext, messaging.WithRabbitMQDialer())

	queues := []messaging.RabbitMQQueue{
		messaging.NewDefaultRabbitMQQueue(connection, appName+"-queue"),
	}
	listener := messaging.NewDefaultMessagingListener[*amqp.Delivery]()
	dispatcher := server.BuildRabbitMQMessageDispatcher(listener, queues...)

	app.Attach("RabbitMQDispatcher", dispatcher)

	if err = app.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
