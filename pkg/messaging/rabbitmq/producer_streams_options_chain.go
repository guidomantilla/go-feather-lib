package rabbitmq

import "github.com/rabbitmq/rabbitmq-stream-go-client/pkg/stream"

type StreamsProducerOptionsChain struct {
	chain []StreamsProducerOptions
}

func StreamsProducerOptionsChainBuilder() *StreamsProducerOptionsChain {
	return &StreamsProducerOptionsChain{
		chain: make([]StreamsProducerOptions, 0),
	}
}

func (options *StreamsProducerOptionsChain) Build() StreamsProducerOptions {
	return func(producer *StreamsProducer) {
		for _, option := range options.chain {
			option(producer)
		}
	}
}

func (options *StreamsProducerOptionsChain) WithProducerOptions(poptions *stream.ProducerOptions) *StreamsProducerOptionsChain {
	options.chain = append(options.chain, streamsProducerOptions.WithProducerOptions(poptions))
	return options
}

func (options *StreamsProducerOptionsChain) WithStreamOptions(soptions *stream.StreamOptions) *StreamsProducerOptionsChain {
	options.chain = append(options.chain, streamsProducerOptions.WithStreamOptions(soptions))
	return options
}
