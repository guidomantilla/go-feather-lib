package messaging

import (
	"crypto/tls"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/stream"
)

func AMQPDialer() ConnectionDialer[*amqp.Connection] {
	return func(url string) (*amqp.Connection, error) {
		return amqp.Dial(url)
	}
}

func AMQPDialerTLS(amqps *tls.Config) ConnectionDialer[*amqp.Connection] {
	return func(url string) (*amqp.Connection, error) {
		return amqp.DialTLS(url, amqps)
	}
}

func StreamsDialer() ConnectionDialer[*stream.Environment] {
	return func(url string) (*stream.Environment, error) {
		return stream.NewEnvironment(stream.NewEnvironmentOptions().SetUri(url))
	}
}

func StreamsDialerTLS(streams *tls.Config) ConnectionDialer[*stream.Environment] {
	return func(url string) (*stream.Environment, error) {
		return stream.NewEnvironment(stream.NewEnvironmentOptions().SetUri(url).SetTLSConfig(streams))
	}
}
