package rabbitmq

import "github.com/rabbitmq/rabbitmq-stream-go-client/pkg/stream"

var streamsProducerOptions_ = NewStreamsProducerOptions()

func NewStreamsProducerOptions() streamsProducerOptions {
	return func(producer *streamsProducer) {
	}
}

type streamsProducerOptions func(*streamsProducer)

func (options streamsProducerOptions) WithProducerOptions(poptions *stream.ProducerOptions) streamsProducerOptions {
	return func(producer *streamsProducer) {
		producer.producerOptions = poptions
	}
}

func (options streamsProducerOptions) WithStreamOptions(soptions *stream.StreamOptions) streamsProducerOptions {
	return func(producer *streamsProducer) {
		producer.streamOptions = soptions
	}
}
