package messaging

var contextOption = NewContextOption()

func NewContextOption() ContextOption {
	return func(context Context) {
	}
}

type ContextOption func(context Context)

func (option ContextOption) WithService(service string) ContextOption {
	return func(context Context) {
		context.set("service", service)
	}
}

func (option ContextOption) WithVhost(vhost string) ContextOption {
	return func(context Context) {
		context.set("vhost", vhost)
	}
}
