package main

import (
	"context"
	"os"

	samqp "github.com/rabbitmq/rabbitmq-stream-go-client/pkg/amqp"

	"github.com/guidomantilla/go-feather-lib/pkg/common/environment"
	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
	cserver "github.com/guidomantilla/go-feather-lib/pkg/common/server"
	"github.com/guidomantilla/go-feather-lib/pkg/common/ssl"
	rabbitmqstreams "github.com/guidomantilla/go-feather-lib/pkg/messaging/rabbitmq/streams"
)

func main() {

	_ = os.Setenv("LOG_LEVEL", "DEBUG")
	cserver.Run("rabbitmq-stream-micro", "1.0.0", func(ctx context.Context, application cserver.Application) error {

		name := "rabbitmq-stream-micro-stream"

		serverName := environment.Value(environment.SslServerName).AsString()
		caCertificate := environment.Value(environment.SslCaCertificate).AsString()
		clientCertificate := environment.Value(environment.SslClientCertificate).AsString()
		clientKey := environment.Value(environment.SslClientKey).AsString()
		tlsConfig, _ := ssl.TLS(serverName, caCertificate, clientCertificate, clientKey)

		messagingContext := rabbitmqstreams.NewContext("rabbitmq-stream+tls://:username::password@:server:vhost",
			"raven-dev", "raven-dev*+", "ubuntu-us-southeast:5551", rabbitmqstreams.ContextOptionBuilder().WithVhost("/").Build())

		{ // Keep an 1:1 relationship between the environment and the consumer

			connection := rabbitmqstreams.NewConnection(messagingContext, rabbitmqstreams.DialerTLS(tlsConfig))
			consumer := rabbitmqstreams.NewConsumer(connection, name)

			application.Attach(rabbitmqstreams.BuildConsumerServer(consumer))
		}

		{ // Keep an 1:1 relationship between the environment and the publisher
			connection := rabbitmqstreams.NewConnection(messagingContext, rabbitmqstreams.DialerTLS(tlsConfig))
			producer := rabbitmqstreams.NewProducer(connection, name)
			if err := producer.Produce(context.Background(), samqp.NewMessage([]byte("Hello, World!"))); err != nil {
				log.Fatal(ctx, "Error producing message: %v", err)
			}
			connection.Close(ctx)
		}

		return nil
	})
}
