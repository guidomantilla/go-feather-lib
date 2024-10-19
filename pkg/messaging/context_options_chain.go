package messaging

type ContextOptionChain struct {
	chain []ContextOption
}

func NewContextOptionChain() *ContextOptionChain {
	return &ContextOptionChain{
		chain: make([]ContextOption, 0),
	}
}

func (options *ContextOptionChain) Build() ContextOption {
	return func(context Context) {
		for _, option := range options.chain {
			option(context)
		}
	}
}

func (options *ContextOptionChain) WithService(service string) *ContextOptionChain {
	options.chain = append(options.chain, contextOption.WithService(service))
	return options
}

func (options *ContextOptionChain) WithVhost(vhost string) *ContextOptionChain {
	options.chain = append(options.chain, contextOption.WithVhost(vhost))
	return options
}
