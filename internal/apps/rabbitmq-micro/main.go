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

	options := []messaging.RabbitMQContextOption{
		messaging.WithFailOver(true),
		//messaging.WithInternalObserver(true),
	}
	messagingContext := messaging.NewDefaultRabbitMQContext("amqp://:username::password@:server/", "raven-dev", "raven-dev*+", "170.187.157.212:5672", options...)
	connection := messaging.NewDefaultRabbitMQConnection(messagingContext)
	defer connection.Close()

	//channel := messaging.NewDefaultRabbitMQChannel(connection)
	//defer channel.Close()

	queue := messaging.NewDefaultRabbitMQQueue(connection, "queue", "consumer")
	defer queue.Close()

	//connection.Connect()
	//channel.Connect()

	go func() {
		log.Info("entering goroutine")
		for {
			log.Info("opening deliveries")
			rabbitChannel, _ := queue.Connect()
			deliveries, _ := rabbitChannel.Consume("queue", "consumer", true, false, false, false, nil)
			for d := range deliveries {
				log.Info(string(d.Body))
			}
			log.Info("closing deliveries")
		}

		log.Info("leaving goroutine")
	}()

	//listener := messaging.NewDefaultRabbitMQQueueMessageListener("my-queue")
	//dispatcher := server.BuildRabbitMQQueueMessageDispatcher(messagingContext, connection, listener)
	//app.Attach("RabbitMQDispatcher", dispatcher)

	if err = app.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
