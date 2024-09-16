package main

import (
	"context"
	"os"
	"syscall"
	"time"

	"github.com/qmdx00/lifecycle"
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/guidomantilla/go-feather-lib/pkg/common/environment"
	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
	"github.com/guidomantilla/go-feather-lib/pkg/common/ssl"
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

	envs := environment.Default()

	serverName := envs.Value(environment.SslServerName).AsString()
	caCertificate := envs.Value(environment.SslCaCertificate).AsString()
	clientCertificate := envs.Value(environment.SslClientCertificate).AsString()
	clientKey := envs.Value(environment.SslClientKey).AsString()
	tlsConfig, _ := ssl.TLS(serverName, caCertificate, clientCertificate, clientKey)

	messagingContext := messaging.NewDefaultMessagingContext("amqps://:username::password@:server:vhost", //?auth_mechanism=EXTERNAL
		"raven-dev", "raven-dev*+", "ubuntu-us-southeast:5671", messaging.WithVhost("/"))

	{ // Keep an 1:1 relationship between the connection, the channel and the consumer
		connection := messaging.NewRabbitMQConnection(messagingContext, messaging.WithRabbitMQDialerTLS(tlsConfig))
		consumer := messaging.NewRabbitMQConsumer(connection, appName+"-queue")

		app.Attach("RabbitMQServer", server.BuildRabbitMQServer(consumer))
	}

	{ // Keep an 1:1 relationship between the connection, the channel and the publisher
		connection := messaging.NewRabbitMQConnection(messagingContext, messaging.WithRabbitMQDialerTLS(tlsConfig))
		producer := messaging.NewRabbitMQProducer(connection, appName+"-queue")

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

	if err = app.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
