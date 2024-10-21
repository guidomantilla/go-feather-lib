package rabbitmq

import (
	"crypto/tls"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/stream"

	"github.com/guidomantilla/go-feather-lib/pkg/messaging"
)

func AMQPDialer() messaging.ConnectionDialer[*amqp.Connection] {
	return func(url string) (*amqp.Connection, error) {
		return amqp.Dial(url)
	}
}

func AMQPDialerTLS(amqps *tls.Config) messaging.ConnectionDialer[*amqp.Connection] {
	return func(url string) (*amqp.Connection, error) {
		return amqp.DialTLS(url, amqps)
	}
}

func StreamsDialer() messaging.ConnectionDialer[*stream.Environment] {
	return func(url string) (*stream.Environment, error) {
		return stream.NewEnvironment(stream.NewEnvironmentOptions().SetUri(url))
	}
}

func StreamsDialerTLS(streams *tls.Config) messaging.ConnectionDialer[*stream.Environment] {
	return func(url string) (*stream.Environment, error) {
		return stream.NewEnvironment(stream.NewEnvironmentOptions().SetUri(url).SetTLSConfig(streams))
	}
}
