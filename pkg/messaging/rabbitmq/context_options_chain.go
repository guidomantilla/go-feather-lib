package rabbitmq

import "github.com/guidomantilla/go-feather-lib/pkg/messaging"

type ContextOptionsChain struct {
	chain []ContextOptions
}

func NewContextOptionChain() *ContextOptionsChain {
	return &ContextOptionsChain{
		chain: make([]ContextOptions, 0),
	}
}

func (options *ContextOptionsChain) Build() ContextOptions {
	return func(context messaging.Context) {
		for _, option := range options.chain {
			option(context)
		}
	}
}

func (options *ContextOptionsChain) WithService(service string) *ContextOptionsChain {
	options.chain = append(options.chain, contextOptions.WithService(service))
	return options
}

func (options *ContextOptionsChain) WithVhost(vhost string) *ContextOptionsChain {
	options.chain = append(options.chain, contextOptions.WithVhost(vhost))
	return options
}
