package rabbitmq

var amqpProducerOptions_ = NewAmqpProducerOptions()

func NewAmqpProducerOptions() amqpProducerOptions {
	return func(producer *amqpProducer) {
	}
}

type amqpProducerOptions func(*amqpProducer)

func (options amqpProducerOptions) WithExchange(exchange string) amqpProducerOptions {
	return func(producer *amqpProducer) {
		producer.exchange = exchange
	}
}

func (options amqpProducerOptions) WithMandatory(mandatory bool) amqpProducerOptions {
	return func(producer *amqpProducer) {
		producer.mandatory = mandatory
	}
}

func (options amqpProducerOptions) WithImmediate(immediate bool) amqpProducerOptions {
	return func(producer *amqpProducer) {
		producer.immediate = immediate
	}
}
