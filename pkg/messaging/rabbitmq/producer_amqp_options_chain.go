package rabbitmq

type amqpProducerOptionsChain struct {
	chain []amqpProducerOptions
}

func AmqpProducerOptionsChainBuilder() *amqpProducerOptionsChain {
	return &amqpProducerOptionsChain{
		chain: make([]amqpProducerOptions, 0),
	}
}

func (options *amqpProducerOptionsChain) Build() amqpProducerOptions {
	return func(producer *amqpProducer) {
		for _, option := range options.chain {
			option(producer)
		}
	}
}

func (options *amqpProducerOptionsChain) WithExchange(exchange string) *amqpProducerOptionsChain {
	options.chain = append(options.chain, amqpProducerOptions_.WithExchange(exchange))
	return options
}

func (options *amqpProducerOptionsChain) WithMandatory(mandatory bool) *amqpProducerOptionsChain {
	options.chain = append(options.chain, amqpProducerOptions_.WithMandatory(mandatory))
	return options
}

func (options *amqpProducerOptionsChain) WithImmediate(immediate bool) *amqpProducerOptionsChain {
	options.chain = append(options.chain, amqpProducerOptions_.WithImmediate(immediate))
	return options
}
