package rabbitmq

type AmqpProducerOptionsChain struct {
	chain []AmqpProducerOptions
}

func AmqpProducerOptionsChainBuilder() *AmqpProducerOptionsChain {
	return &AmqpProducerOptionsChain{
		chain: make([]AmqpProducerOptions, 0),
	}
}

func (options *AmqpProducerOptionsChain) Build() AmqpProducerOptions {
	return func(producer *AmqpProducer) {
		for _, option := range options.chain {
			option(producer)
		}
	}
}

func (options *AmqpProducerOptionsChain) WithExchange(exchange string) *AmqpProducerOptionsChain {
	options.chain = append(options.chain, amqpProducerOptions.WithExchange(exchange))
	return options
}

func (options *AmqpProducerOptionsChain) WithMandatory(mandatory bool) *AmqpProducerOptionsChain {
	options.chain = append(options.chain, amqpProducerOptions.WithMandatory(mandatory))
	return options
}

func (options *AmqpProducerOptionsChain) WithImmediate(immediate bool) *AmqpProducerOptionsChain {
	options.chain = append(options.chain, amqpProducerOptions.WithImmediate(immediate))
	return options
}
