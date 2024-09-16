package main

import (
	"os"
	"syscall"

	"github.com/qmdx00/lifecycle"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
	"github.com/guidomantilla/go-feather-lib/pkg/common/ssl"
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
	tlsConfig, _ := ssl.BuildTLS()
	messagingContext := messaging.NewDefaultMessagingContext("rabbitmq-stream+tls://:username::password@:server:vhost",
		"raven-dev", "raven-dev*+", "ubuntu-us-southeast:5551", messaging.WithVhost("/"))
	connection := messaging.NewRabbitMQConnection(messagingContext, messaging.WithRabbitMQStreamsDialerTLS(tlsConfig))

	consumer := messaging.NewRabbitMQStreamsConsumer(connection, "rabbitmq-stream-micro-stream")
	app.Attach("RabbitMQServer", server.BuildRabbitMQServer(consumer))

	if err = app.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
