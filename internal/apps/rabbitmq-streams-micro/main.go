package main

import (
	"os"
	"syscall"

	"github.com/qmdx00/lifecycle"
	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/amqp"
	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/stream"

	"github.com/guidomantilla/go-feather-lib/pkg/common/environment"
	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
	"github.com/guidomantilla/go-feather-lib/pkg/common/ssl"
	"github.com/guidomantilla/go-feather-lib/pkg/messaging"
	"github.com/guidomantilla/go-feather-lib/pkg/server"
)

func main() {

	var err error
	appName, version := "rabbitmq-stream-micro", "1.0.0"
	os.Setenv("LOG_LEVEL", "DEBUG")
	log.Custom()

	app := lifecycle.NewApp(
		lifecycle.WithName(appName), lifecycle.WithVersion(version),
		lifecycle.WithSignal(syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGKILL),
	)

	envs := environment.Default()

	serverName := envs.Value(environment.SslServerName).AsString()
	caCertificate := envs.Value(environment.SslCaCertificate).AsString()
	clientCertificate := envs.Value(environment.SslClientCertificate).AsString()
	clientKey := envs.Value(environment.SslClientKey).AsString()
	tlsConfig, _ := ssl.TLS(serverName, caCertificate, clientCertificate, clientKey)

	messagingContext := messaging.NewDefaultMessagingContext("rabbitmq-stream+tls://:username::password@:server:vhost",
		"raven-dev", "raven-dev*+", "ubuntu-us-southeast:5551", messaging.WithVhost("/"))

	{ // Keep an 1:1 relationship between the environment and the consumer
		connection := messaging.NewRabbitMQConnection(messagingContext, messaging.WithRabbitMQStreamsDialerTLS(tlsConfig))
		consumer := messaging.NewRabbitMQStreamsConsumer(connection, appName+"-stream")

		app.Attach("RabbitMQServer", server.BuildRabbitMQServer(consumer))
	}

	{ // Keep an 1:1 relationship between the environment and the publisher
		connection := messaging.NewRabbitMQConnection(messagingContext, messaging.WithRabbitMQStreamsDialerTLS(tlsConfig))
		conn, _ := connection.Connect()
		producer, _ := conn.NewProducer(appName+"-stream", stream.NewProducerOptions())
		_ = producer.Send(amqp.NewMessage([]byte("Hello world")))
		connection.Close()
	}

	if err = app.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
