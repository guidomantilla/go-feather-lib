package main

import (
	"context"
	"os"

	samqp "github.com/rabbitmq/rabbitmq-stream-go-client/pkg/amqp"

	"github.com/guidomantilla/go-feather-lib/pkg/common/environment"
	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
	cserver "github.com/guidomantilla/go-feather-lib/pkg/common/server"
	"github.com/guidomantilla/go-feather-lib/pkg/common/ssl"
	"github.com/guidomantilla/go-feather-lib/pkg/messaging/rabbitmq"
	"github.com/guidomantilla/go-feather-lib/pkg/server"
)

func main() {

	_ = os.Setenv("LOG_LEVEL", "DEBUG")
	cserver.Run("rabbitmq-stream-micro", "1.0.0", func(application cserver.Application) error {

		name := "rabbitmq-stream-micro-stream"

		serverName := environment.Value(environment.SslServerName).AsString()
		caCertificate := environment.Value(environment.SslCaCertificate).AsString()
		clientCertificate := environment.Value(environment.SslClientCertificate).AsString()
		clientKey := environment.Value(environment.SslClientKey).AsString()
		tlsConfig, _ := ssl.TLS(serverName, caCertificate, clientCertificate, clientKey)

		messagingContext := rabbitmq.NewContext("rabbitmq-stream+tls://:username::password@:server:vhost",
			"raven-dev", "raven-dev*+", "ubuntu-us-southeast:5551", rabbitmq.NewContextOptionChain().WithVhost("/").Build())

		{ // Keep an 1:1 relationship between the environment and the consumer

			connection := rabbitmq.NewConnection(messagingContext, rabbitmq.StreamsDialerTLS(tlsConfig))
			consumer := rabbitmq.NewStreamsConsumer(connection, name)

			application.Attach(server.BuildRabbitMQServer(consumer))
		}

		{ // Keep an 1:1 relationship between the environment and the publisher
			connection := rabbitmq.NewConnection(messagingContext, rabbitmq.StreamsDialerTLS(tlsConfig))
			producer := rabbitmq.NewStreamsProducer(connection, name)
			if err := producer.Produce(context.Background(), samqp.NewMessage([]byte("Hello, World!"))); err != nil {
				log.Fatal("Error producing message: %v", err)
			}
			connection.Close()
		}

		return nil
	})
}
