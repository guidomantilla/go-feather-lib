package streams

import "github.com/rabbitmq/rabbitmq-stream-go-client/pkg/stream"

type producerOptionsChain struct {
	chain []ProducerOptions
}

func ProducerOptionsChainBuilder() *producerOptionsChain {
	return &producerOptionsChain{
		chain: make([]ProducerOptions, 0),
	}
}

func (options *producerOptionsChain) Build() ProducerOptions {
	return func(producer Producer) {
		for _, option := range options.chain {
			option(producer)
		}
	}
}

func (options *producerOptionsChain) WithProducerOptions(poptions *stream.ProducerOptions) *producerOptionsChain {
	options.chain = append(options.chain, producerOptions.WithProducerOptions(poptions))
	return options
}

func (options *producerOptionsChain) WithStreamOptions(soptions *stream.StreamOptions) *producerOptionsChain {
	options.chain = append(options.chain, producerOptions.WithStreamOptions(soptions))
	return options
}
