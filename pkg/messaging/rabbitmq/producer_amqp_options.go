package rabbitmq

var amqpProducerOptions = NewAmqpProducerOptions()

func NewAmqpProducerOptions() AmqpProducerOptions {
	return func(producer *AmqpProducer) {
	}
}

type AmqpProducerOptions func(*AmqpProducer)

func (options AmqpProducerOptions) WithExchange(exchange string) AmqpProducerOptions {
	return func(producer *AmqpProducer) {
		producer.exchange = exchange
	}
}

func (options AmqpProducerOptions) WithMandatory(mandatory bool) AmqpProducerOptions {
	return func(producer *AmqpProducer) {
		producer.mandatory = mandatory
	}
}

func (options AmqpProducerOptions) WithImmediate(immediate bool) AmqpProducerOptions {
	return func(producer *AmqpProducer) {
		producer.immediate = immediate
	}
}
