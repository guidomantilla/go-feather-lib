package amqp

import (
	"crypto/tls"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Dialer() ConnectionDialer {
	return func(url string) (*amqp.Connection, error) {
		return amqp.Dial(url)
	}
}

func DialerTLS(amqps *tls.Config) ConnectionDialer {
	return func(url string) (*amqp.Connection, error) {
		return amqp.DialTLS(url, amqps)
	}
}
