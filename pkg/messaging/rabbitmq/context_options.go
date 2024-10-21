package rabbitmq

import "github.com/guidomantilla/go-feather-lib/pkg/messaging"

var contextOptions = NewContextOption()

func NewContextOption() ContextOptions {
	return func(context messaging.Context) {
	}
}

type ContextOptions func(context messaging.Context)

func (option ContextOptions) WithService(service string) ContextOptions {
	return func(context messaging.Context) {
		context.Set("service", service)
	}
}

func (option ContextOptions) WithVhost(vhost string) ContextOptions {
	return func(context messaging.Context) {
		context.Set("vhost", vhost)
	}
}
