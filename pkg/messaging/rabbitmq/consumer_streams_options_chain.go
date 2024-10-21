package rabbitmq

import (
	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/amqp"
	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/stream"

	"github.com/guidomantilla/go-feather-lib/pkg/messaging"
)

type StreamsConsumerOptionsChain struct {
	chain []StreamsConsumerOptions
}

func StreamsConsumerOptionsChainBuilder() *StreamsConsumerOptionsChain {
	return &StreamsConsumerOptionsChain{
		chain: make([]StreamsConsumerOptions, 0),
	}
}

func (options *StreamsConsumerOptionsChain) Build() StreamsConsumerOptions {
	return func(consumer *StreamsConsumer) {
		for _, option := range options.chain {
			option(consumer)
		}
	}
}

func (options *StreamsConsumerOptionsChain) WithStreamOptions(soptions *stream.StreamOptions) *StreamsConsumerOptionsChain {
	options.chain = append(options.chain, streamsConsumerOptions.WithStreamOptions(soptions))
	return options
}

func (options *StreamsConsumerOptionsChain) WithConsumerOptions(coptions *stream.ConsumerOptions) *StreamsConsumerOptionsChain {
	options.chain = append(options.chain, streamsConsumerOptions.WithConsumerOptions(coptions))
	return options
}

func (options *StreamsConsumerOptionsChain) WithStreamsListener(listener messaging.Listener[*amqp.Message]) *StreamsConsumerOptionsChain {
	options.chain = append(options.chain, streamsConsumerOptions.WithStreamsListener(listener))
	return options
}
