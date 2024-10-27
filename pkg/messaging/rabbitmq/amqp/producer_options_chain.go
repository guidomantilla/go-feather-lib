package amqp

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

func (options *producerOptionsChain) WithExchange(exchange string) *producerOptionsChain {
	options.chain = append(options.chain, producerOptions.WithExchange(exchange))
	return options
}

func (options *producerOptionsChain) WithMandatory(mandatory bool) *producerOptionsChain {
	options.chain = append(options.chain, producerOptions.WithMandatory(mandatory))
	return options
}

func (options *producerOptionsChain) WithImmediate(immediate bool) *producerOptionsChain {
	options.chain = append(options.chain, producerOptions.WithImmediate(immediate))
	return options
}
