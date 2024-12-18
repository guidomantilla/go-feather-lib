package main

import (
	"context"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/guidomantilla/go-feather-lib/pkg/common/environment"
	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
	cserver "github.com/guidomantilla/go-feather-lib/pkg/common/server"
	"github.com/guidomantilla/go-feather-lib/pkg/common/ssl"
	"github.com/guidomantilla/go-feather-lib/pkg/messaging/rabbitmq"
	"github.com/guidomantilla/go-feather-lib/pkg/server"
)

func main() {

	_ = os.Setenv("LOG_LEVEL", "DEBUG")
	cserver.Run("rabbitmq-micro", "1.0.0", func(application cserver.Application) error {

		name := "rabbitmq-micro-queue"

		serverName := environment.Value(environment.SslServerName).AsString()
		caCertificate := environment.Value(environment.SslCaCertificate).AsString()
		clientCertificate := environment.Value(environment.SslClientCertificate).AsString()
		clientKey := environment.Value(environment.SslClientKey).AsString()
		tlsConfig, _ := ssl.TLS(serverName, caCertificate, clientCertificate, clientKey)

		messagingContext := rabbitmq.NewContext("amqps://:username::password@:server:vhost", //?auth_mechanism=EXTERNAL
			"raven-dev", "raven-dev*+", "ubuntu-us-southeast:5671", rabbitmq.NewContextOptionChain().WithVhost("/").Build())

		{ // Keep an 1:1 relationship between the connection, the channel and the consumer

			connection := rabbitmq.NewConnection(messagingContext, rabbitmq.AMQPDialerTLS(tlsConfig))
			consumer := rabbitmq.NewAmqpConsumer(connection, name)

			application.Attach(server.BuildRabbitMQServer(consumer))
		}

		{ // Keep an 1:1 relationship between the connection, the channel and the publisher
			connection := rabbitmq.NewConnection(messagingContext, rabbitmq.AMQPDialerTLS(tlsConfig))
			producer := rabbitmq.NewAmqpProducer(connection, name)

			if err := producer.Produce(context.Background(), &amqp.Publishing{
				Headers:         nil,
				ContentType:     "",
				ContentEncoding: "",
				DeliveryMode:    0,
				Priority:        0,
				CorrelationId:   "",
				ReplyTo:         "",
				Expiration:      "",
				MessageId:       "",
				Timestamp:       time.Time{},
				Type:            "",
				UserId:          "",
				AppId:           "",
				Body:            []byte("Hello, World! xxx"),
			}); err != nil {
				log.Fatal("Error producing message: %v", err)
			}

			producer.Close()
		}

		return nil
	})
}
