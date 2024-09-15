package main

import (
	"fmt"
	"os"
	"syscall"

	"github.com/qmdx00/lifecycle"
	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/amqp"
	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/stream"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
	"github.com/guidomantilla/go-feather-lib/pkg/messaging"
	"github.com/guidomantilla/go-feather-lib/pkg/server"
)

func CheckErrReceive(err error) {
	if err != nil {
		fmt.Printf("%s ", err)
		os.Exit(1)
	}
}

func main() {

	var err error
	appName, version := "rabbitmq-stream-micro-stream", "1.0.0"
	os.Setenv("LOG_LEVEL", "DEBUG")
	log.Custom()

	messagingContext := messaging.NewDefaultMessagingContext("rabbitmq-stream://:username::password@:server:vhost",
		"raven-dev", "raven-dev*+", "170.187.157.212:5552", messaging.WithVhost("/"))

	connection := messaging.NewRabbitMQConnection(messagingContext, messaging.WithRabbitMQStreamsDialer())
	streams := messaging.NewDefaultRabbitMQStreams(connection, "rabbitmq-stream-micro-stream")
	options := stream.NewConsumerOptions().SetOffset(stream.OffsetSpecification{}.First()).SetConsumerName("rabbitmq-stream-micro-stream")

	messagesHandler := func(consumerContext stream.ConsumerContext, message *amqp.Message) {
		fmt.Printf("Stream: %s - Received message: %s\n", consumerContext.Consumer.GetStreamName(), message.Data)
	}

	go func() {
		for {
			select {
			default:
				var err error

				var env *stream.Environment
				if env, err = streams.Connect(); err != nil {
					log.Error(fmt.Sprintf("rabbitmq dispatcher - error: %s", err.Error()))
					continue
				}

				var consumer *stream.Consumer
				if consumer, err = env.NewConsumer("rabbitmq-stream-micro-stream", messagesHandler, options); err != nil {
					log.Error(fmt.Sprintf("rabbitmq dispatcher - error: %s", err.Error()))
					continue
				}

				for event := range consumer.NotifyClose() {
					log.Info("Stream: %s - Consumer closed", "rabbitmq-stream-micro-stream", event.Reason)
				}
			}
		}
	}()

	app := lifecycle.NewApp(
		lifecycle.WithName(appName), lifecycle.WithVersion(version),
		lifecycle.WithSignal(syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGKILL),
	)

	app.Attach("DummyServer", server.BuildDummyServer())

	if err = app.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
