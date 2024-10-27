package amqp

type contextOptionsChain struct {
	chain []ContextOptions
}

func ContextOptionBuilder() *contextOptionsChain {
	return &contextOptionsChain{
		chain: make([]ContextOptions, 0),
	}
}

func (options *contextOptionsChain) Build() ContextOptions {
	return func(context Context) {
		for _, option := range options.chain {
			option(context)
		}
	}
}

func (options *contextOptionsChain) WithService(service string) *contextOptionsChain {
	options.chain = append(options.chain, contextOptions.WithService(service))
	return options
}

func (options *contextOptionsChain) WithVhost(vhost string) *contextOptionsChain {
	options.chain = append(options.chain, contextOptions.WithVhost(vhost))
	return options
}
