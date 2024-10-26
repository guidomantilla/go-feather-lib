package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type AmqpConsumerOptionsChain struct {
	chain []AmqpConsumerOptions
}

func AmqpConsumerOptionsChainBuilder() *AmqpConsumerOptionsChain {
	return &AmqpConsumerOptionsChain{
		chain: make([]AmqpConsumerOptions, 0),
	}
}

func (options *AmqpConsumerOptionsChain) Build() AmqpConsumerOptions {
	return func(consumer Consumer) {
		for _, option := range options.chain {
			option(consumer)
		}
	}
}

func (options *AmqpConsumerOptionsChain) WithRabbitMQListener(listener Listener[*amqp.Delivery]) *AmqpConsumerOptionsChain {
	options.chain = append(options.chain, amqpConsumerOptions.WithRabbitMQListener(listener))
	return options
}

func (options *AmqpConsumerOptionsChain) WithAutoAck(autoAck bool) *AmqpConsumerOptionsChain {
	options.chain = append(options.chain, amqpConsumerOptions.WithAutoAck(autoAck))
	return options
}

func (options *AmqpConsumerOptionsChain) WithNoLocal(noLocal bool) *AmqpConsumerOptionsChain {
	options.chain = append(options.chain, amqpConsumerOptions.WithNoLocal(noLocal))
	return options
}

func (options *AmqpConsumerOptionsChain) WithDurable(durable bool) *AmqpConsumerOptionsChain {
	options.chain = append(options.chain, amqpConsumerOptions.WithDurable(durable))
	return options
}

func (options *AmqpConsumerOptionsChain) WithAutoDelete(autoDelete bool) *AmqpConsumerOptionsChain {
	options.chain = append(options.chain, amqpConsumerOptions.WithAutoDelete(autoDelete))
	return options
}

func (options *AmqpConsumerOptionsChain) WithExclusive(exclusive bool) *AmqpConsumerOptionsChain {
	options.chain = append(options.chain, amqpConsumerOptions.WithExclusive(exclusive))
	return options
}

func (options *AmqpConsumerOptionsChain) WithNoWait(noWait bool) *AmqpConsumerOptionsChain {
	options.chain = append(options.chain, amqpConsumerOptions.WithNoWait(noWait))
	return options
}

func (options *AmqpConsumerOptionsChain) WithArgs(args amqp.Table) *AmqpConsumerOptionsChain {
	options.chain = append(options.chain, amqpConsumerOptions.WithArgs(args))
	return options
}
