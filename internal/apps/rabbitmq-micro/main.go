package main

import (
	"fmt"
	"os"
	"syscall"

	"github.com/qmdx00/lifecycle"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
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
	app.Attach("DummyServer", server.BuildDummyServer())

	deliveries, err := queue.Consume()
	if err != nil {
		log.Fatal(err.Error())
	}

	go func() {
		for {

			select {
			case deliveryChan := <-deliveries:
				for {
					message, ok := <-deliveryChan
					if !ok {
						break
					}
					log.Info(string(message.Body))
				}
			}
		}
	}()

	//listener := messaging.NewDefaultRabbitMQQueueMessageListener("my-queue")
	//dispatcher := server.BuildRabbitMQQueueMessageDispatcher(messagingContext, connection, listener)
	//app.Attach("RabbitMQDispatcher", dispatcher)

	if err = app.Run(); err != nil {
		log.Fatal(err.Error())
	}

	if err2 := recover(); err2 != nil {
		log.Fatal(fmt.Sprintf("panic: %v", err2))
	}
}
