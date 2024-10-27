package streams

import (
	"crypto/tls"

	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/stream"
)

func Dialer() ConnectionDialer {
	return func(url string) (*stream.Environment, error) {
		return stream.NewEnvironment(stream.NewEnvironmentOptions().SetUri(url))
	}
}

func DialerTLS(streams *tls.Config) ConnectionDialer {
	return func(url string) (*stream.Environment, error) {
		return stream.NewEnvironment(stream.NewEnvironmentOptions().SetUri(url).SetTLSConfig(streams))
	}
}
