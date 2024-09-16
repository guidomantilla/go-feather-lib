package main

import (
	"github.com/guidomantilla/go-feather-lib/pkg/ssl"
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

	tlsConfig, _ := ssl.BuildTLS()
	messagingContext := messaging.NewDefaultMessagingContext("amqps://:username::password@:server:vhost", //?auth_mechanism=EXTERNAL
		"raven-dev", "raven-dev*+", "ubuntu-us-southeast:5671", messaging.WithVhost("/"))
	connection := messaging.NewRabbitMQConnection(messagingContext, messaging.WithRabbitMQDialerTLS(tlsConfig))

	consumer := messaging.NewRabbitMQConsumer(connection, appName+"-queue")
	app.Attach("RabbitMQServer", server.BuildRabbitMQServer(consumer))

	if err = app.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
