package streams

import (
	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/stream"
)

type consumerOptionsChain struct {
	chain []ConsumerOptions
}

func ConsumerOptionsBuilder() *consumerOptionsChain {
	return &consumerOptionsChain{
		chain: make([]ConsumerOptions, 0),
	}
}

func (options *consumerOptionsChain) Build() ConsumerOptions {
	return func(consumer *consumer) {
		for _, option := range options.chain {
			option(consumer)
		}
	}
}

func (options *consumerOptionsChain) WithStreamOptions(soptions *stream.StreamOptions) *consumerOptionsChain {
	options.chain = append(options.chain, consumerOptions.WithStreamOptions(soptions))
	return options
}

func (options *consumerOptionsChain) WithConsumerOptions(coptions *stream.ConsumerOptions) *consumerOptionsChain {
	options.chain = append(options.chain, consumerOptions.WithConsumerOptions(coptions))
	return options
}

func (options *consumerOptionsChain) WithListener(listener Listener) *consumerOptionsChain {
	options.chain = append(options.chain, consumerOptions.WithListener(listener))
	return options
}
