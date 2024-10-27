package amqp

var producerOptions = NewProducerOptions()

func NewProducerOptions() ProducerOptions {
	return func(producer Producer) {
	}
}

type ProducerOptions func(producer Producer)

func (options ProducerOptions) WithExchange(exchange string) ProducerOptions {
	return func(producer Producer) {
		//producer.exchange = exchange
	}
}

func (options ProducerOptions) WithMandatory(mandatory bool) ProducerOptions {
	return func(producer Producer) {
		//producer.mandatory = mandatory
	}
}

func (options ProducerOptions) WithImmediate(immediate bool) ProducerOptions {
	return func(producer Producer) {
		//producer.immediate = immediate
	}
}
