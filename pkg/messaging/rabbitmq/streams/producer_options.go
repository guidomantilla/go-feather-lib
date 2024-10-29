package streams

import "github.com/rabbitmq/rabbitmq-stream-go-client/pkg/stream"

var producerOptions = NewProducerOptions()

func NewProducerOptions() ProducerOptions {
	return func(producer Producer) {
	}
}

type ProducerOptions func(producer Producer)

func (options ProducerOptions) WithProducerOptions(poptions *stream.ProducerOptions) ProducerOptions {
	return func(producer Producer) {
		producer.Set("producerOptions", poptions)
	}
}

func (options ProducerOptions) WithStreamOptions(soptions *stream.StreamOptions) ProducerOptions {
	return func(producer Producer) {
		producer.Set("streamOptions", soptions)
	}
}
