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

	stopCh := make(chan struct{})
	go func() {
		for {
			select {
			case <-stopCh:
				return
			default:

				connection := messaging.NewRabbitMQConnection(messagingContext, messaging.WithRabbitMQStreamsDialer())
				env, err := connection.Connect()
				if err != nil {
					log.Error(fmt.Sprintf("rabbitmq dispatcher - error: %s", err.Error()))
					continue
				}

				streamName := "rabbitmq-stream-micro-stream"
				err = env.DeclareStream(streamName, &stream.StreamOptions{MaxLengthBytes: stream.ByteCapacity{}.GB(2)})
				if err != nil {
					log.Error(fmt.Sprintf("rabbitmq dispatcher - error: %s", err.Error()))
					continue
				}

				messagesHandler := func(consumerContext stream.ConsumerContext, message *amqp.Message) {
					fmt.Printf("Stream: %s - Received message: %s\n", consumerContext.Consumer.GetStreamName(), message.Data)
				}

				consumer, err := env.NewConsumer(streamName, messagesHandler, stream.NewConsumerOptions().SetOffset(stream.OffsetSpecification{}.First()))
				if err != nil {
					log.Error(fmt.Sprintf("rabbitmq dispatcher - error: %s", err.Error()))
					continue
				}

				onClose := consumer.NotifyClose()
				for _ = range onClose {
					log.Info("Stream: %s - Consumer closed", streamName)
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
