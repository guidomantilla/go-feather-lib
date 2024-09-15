package messaging

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/stream"
)

var RabbitMQDialer MessagingConnectionDialer[*amqp.Connection] = amqp.Dial
var RabbitMQStreamsDialer MessagingConnectionDialer[*stream.Environment] = delegateRabbitMQStreamsDialer

func delegateRabbitMQStreamsDialer(url string) (*stream.Environment, error) {
	return stream.NewEnvironment(stream.NewEnvironmentOptions().SetUri(url))
}
