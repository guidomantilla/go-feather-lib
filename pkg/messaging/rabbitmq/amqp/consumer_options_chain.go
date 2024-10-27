package amqp

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type consumerOptionsChain struct {
	chain []ConsumerOptions
}

func ConsumerOptionsChainBuilder() *consumerOptionsChain {
	return &consumerOptionsChain{
		chain: make([]ConsumerOptions, 0),
	}
}

func (options *consumerOptionsChain) Build() ConsumerOptions {
	return func(consumer Consumer) {
		for _, option := range options.chain {
			option(consumer)
		}
	}
}

func (options *consumerOptionsChain) WithListener(listener Listener) *consumerOptionsChain {
	options.chain = append(options.chain, consumerOptions.WithListener(listener))
	return options
}

func (options *consumerOptionsChain) WithAutoAck(autoAck bool) *consumerOptionsChain {
	options.chain = append(options.chain, consumerOptions.WithAutoAck(autoAck))
	return options
}

func (options *consumerOptionsChain) WithNoLocal(noLocal bool) *consumerOptionsChain {
	options.chain = append(options.chain, consumerOptions.WithNoLocal(noLocal))
	return options
}

func (options *consumerOptionsChain) WithDurable(durable bool) *consumerOptionsChain {
	options.chain = append(options.chain, consumerOptions.WithDurable(durable))
	return options
}

func (options *consumerOptionsChain) WithAutoDelete(autoDelete bool) *consumerOptionsChain {
	options.chain = append(options.chain, consumerOptions.WithAutoDelete(autoDelete))
	return options
}

func (options *consumerOptionsChain) WithExclusive(exclusive bool) *consumerOptionsChain {
	options.chain = append(options.chain, consumerOptions.WithExclusive(exclusive))
	return options
}

func (options *consumerOptionsChain) WithNoWait(noWait bool) *consumerOptionsChain {
	options.chain = append(options.chain, consumerOptions.WithNoWait(noWait))
	return options
}

func (options *consumerOptionsChain) WithArgs(args amqp.Table) *consumerOptionsChain {
	options.chain = append(options.chain, consumerOptions.WithArgs(args))
	return options
}
