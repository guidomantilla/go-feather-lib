package rabbitmq

import "github.com/rabbitmq/rabbitmq-stream-go-client/pkg/stream"

type streamsProducerOptionsChain struct {
	chain []streamsProducerOptions
}

func StreamsProducerOptionsChainBuilder() *streamsProducerOptionsChain {
	return &streamsProducerOptionsChain{
		chain: make([]streamsProducerOptions, 0),
	}
}

func (options *streamsProducerOptionsChain) Build() streamsProducerOptions {
	return func(producer *streamsProducer) {
		for _, option := range options.chain {
			option(producer)
		}
	}
}

func (options *streamsProducerOptionsChain) WithProducerOptions(poptions *stream.ProducerOptions) *streamsProducerOptionsChain {
	options.chain = append(options.chain, streamsProducerOptions_.WithProducerOptions(poptions))
	return options
}

func (options *streamsProducerOptionsChain) WithStreamOptions(soptions *stream.StreamOptions) *streamsProducerOptionsChain {
	options.chain = append(options.chain, streamsProducerOptions_.WithStreamOptions(soptions))
	return options
}
