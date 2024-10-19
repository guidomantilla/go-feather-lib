package messaging

import amqp "github.com/rabbitmq/amqp091-go"

type RabbitMQConsumerOptionsChain struct {
	chain []RabbitMQConsumerOptions
}

func RabbitMQConsumerOptionsChainBuilder() *RabbitMQConsumerOptionsChain {
	return &RabbitMQConsumerOptionsChain{
		chain: make([]RabbitMQConsumerOptions, 0),
	}
}

func (options *RabbitMQConsumerOptionsChain) Build() RabbitMQConsumerOptions {
	return func(consumer Consumer) {
		for _, option := range options.chain {
			option(consumer)
		}
	}
}

func (options *RabbitMQConsumerOptionsChain) WithRabbitMQListener(listener Listener[*amqp.Delivery]) *RabbitMQConsumerOptionsChain {
	options.chain = append(options.chain, rabbitMQConsumerOptions.WithRabbitMQListener(listener))
	return options
}

func (options *RabbitMQConsumerOptionsChain) WithAutoAck(autoAck bool) *RabbitMQConsumerOptionsChain {
	options.chain = append(options.chain, rabbitMQConsumerOptions.WithAutoAck(autoAck))
	return options
}

func (options *RabbitMQConsumerOptionsChain) WithNoLocal(noLocal bool) *RabbitMQConsumerOptionsChain {
	options.chain = append(options.chain, rabbitMQConsumerOptions.WithNoLocal(noLocal))
	return options
}

func (options *RabbitMQConsumerOptionsChain) WithDurable(durable bool) *RabbitMQConsumerOptionsChain {
	options.chain = append(options.chain, rabbitMQConsumerOptions.WithDurable(durable))
	return options
}

func (options *RabbitMQConsumerOptionsChain) WithAutoDelete(autoDelete bool) *RabbitMQConsumerOptionsChain {
	options.chain = append(options.chain, rabbitMQConsumerOptions.WithAutoDelete(autoDelete))
	return options
}

func (options *RabbitMQConsumerOptionsChain) WithExclusive(exclusive bool) *RabbitMQConsumerOptionsChain {
	options.chain = append(options.chain, rabbitMQConsumerOptions.WithExclusive(exclusive))
	return options
}

func (options *RabbitMQConsumerOptionsChain) WithNoWait(noWait bool) *RabbitMQConsumerOptionsChain {
	options.chain = append(options.chain, rabbitMQConsumerOptions.WithNoWait(noWait))
	return options
}

func (options *RabbitMQConsumerOptionsChain) WithArgs(args amqp.Table) *RabbitMQConsumerOptionsChain {
	options.chain = append(options.chain, rabbitMQConsumerOptions.WithArgs(args))
	return options
}
