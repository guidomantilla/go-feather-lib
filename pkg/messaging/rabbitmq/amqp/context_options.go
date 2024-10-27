package amqp

var contextOptions = NewContextOption()

func NewContextOption() ContextOptions {
	return func(context Context) {
	}
}

type ContextOptions func(context Context)

func (option ContextOptions) WithService(service string) ContextOptions {
	return func(context Context) {
		context.Set("service", service)
	}
}

func (option ContextOptions) WithVhost(vhost string) ContextOptions {
	return func(context Context) {
		context.Set("vhost", vhost)
	}
}
