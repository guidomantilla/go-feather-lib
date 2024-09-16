package main

import (
	"fmt"
	"os"
	"syscall"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
	"github.com/guidomantilla/go-feather-lib/pkg/messaging"
	"github.com/guidomantilla/go-feather-lib/pkg/server"
	"github.com/qmdx00/lifecycle"
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
	streams := messaging.NewRabbitMQStreams(connection, "rabbitmq-stream-micro-stream")

	go func() {
		for {
			select {
			default:
				var err error
				var closeChannel chan string
				if closeChannel, err = streams.Consume(); err != nil {
					log.Error(fmt.Sprintf("rabbitmq dispatcher - error: %s", err.Error()))
					continue
				}
				<-closeChannel
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
