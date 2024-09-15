package main

import (
	"github.com/guidomantilla/go-feather-lib/pkg/server"
	"os"
	"syscall"

	"github.com/qmdx00/lifecycle"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
	"github.com/guidomantilla/go-feather-lib/pkg/messaging"
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
	app.Attach("DummyServer", server.BuildDummyServer())

	messagingContext := messaging.NewDefaultRabbitMQContext("amqp://:username::password@:server/", "raven-dev", "raven-dev*+", "170.187.157.212:5672")
	connection := messaging.NewDefaultRabbitMQConnection(messagingContext)
	defer connection.Close()

	go func() {
		log.Info("entering goroutine - queue")

		for {
			log.Info("opening deliveries - queue")

			queue := messaging.NewDefaultRabbitMQQueue(connection, "queue", "consumer-queue")
			//defer queue.Close()

			rabbitChannel, _ := queue.Connect()
			deliveries, _ := rabbitChannel.Consume("queue", "consumer-queue", true, false, false, false, nil)
			for d := range deliveries {
				log.Info(string(d.Body))
			}
			log.Info("closing deliveries - queue")
		}
		log.Info("leaving goroutine - queue")
	}()

	go func() {
		log.Info("entering goroutine - my-queue")

		for {
			log.Info("opening deliveries - my-queue")

			queue := messaging.NewDefaultRabbitMQQueue(connection, "my-queue", "consumer-my-queue")
			//defer queue.Close()

			rabbitChannel, _ := queue.Connect()
			deliveries, _ := rabbitChannel.Consume("my-queue", "consumer-my-queue", true, false, false, false, nil)
			for d := range deliveries {
				log.Info(string(d.Body))
			}
			log.Info("closing deliveries - my-queue")
		}
		log.Info("leaving goroutine - my-queue")
	}()

	//listener := messaging.NewDefaultRabbitMQQueueMessageListener("my-queue")
	//dispatcher := server.BuildRabbitMQQueueMessageDispatcher(messagingContext, connection, listener)
	//app.Attach("RabbitMQDispatcher", dispatcher)

	if err = app.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
