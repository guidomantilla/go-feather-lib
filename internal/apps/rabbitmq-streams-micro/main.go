package main

import (
	"bufio"
	"fmt"
	"github.com/guidomantilla/go-feather-lib/pkg/messaging"
	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/amqp"
	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/stream"
	"os"
	"syscall"

	"github.com/qmdx00/lifecycle"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
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

	messagingContext := messaging.NewDefaultRabbitMQContext("rabbitmq-stream://:username::password@:server:vhost", "raven-dev", "raven-dev*+", "170.187.157.212:5552", "/")
	connection := messaging.NewDefaultRabbitMQStreamsConnection(messagingContext)
	env, err := connection.Connect()

	CheckErrReceive(err)

	streamName := "rabbitmq-stream-micro-stream"
	err = env.DeclareStream(streamName, &stream.StreamOptions{MaxLengthBytes: stream.ByteCapacity{}.GB(2)})
	CheckErrReceive(err)

	messagesHandler := func(consumerContext stream.ConsumerContext, message *amqp.Message) {
		fmt.Printf("Stream: %s - Received message: %s\n", consumerContext.Consumer.GetStreamName(), message.Data)
	}

	consumer, err := env.NewConsumer(streamName, messagesHandler, stream.NewConsumerOptions().SetOffset(stream.OffsetSpecification{}.First()))
	CheckErrReceive(err)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println(" [x] Waiting for messages. enter to close the consumer")
	_, _ = reader.ReadString('\n')
	err = consumer.Close()
	CheckErrReceive(err)

	app := lifecycle.NewApp(
		lifecycle.WithName(appName), lifecycle.WithVersion(version),
		lifecycle.WithSignal(syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGKILL),
	)

	app.Attach("DummyServer", server.BuildDummyServer())

	if err = app.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
