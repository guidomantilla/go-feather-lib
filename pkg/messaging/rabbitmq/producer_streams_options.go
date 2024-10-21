package rabbitmq

import "github.com/rabbitmq/rabbitmq-stream-go-client/pkg/stream"

var streamsProducerOptions = NewStreamsProducerOptions()

func NewStreamsProducerOptions() StreamsProducerOptions {
	return func(producer *StreamsProducer) {
	}
}

type StreamsProducerOptions func(*StreamsProducer)

func (options StreamsProducerOptions) WithProducerOptions(poptions *stream.ProducerOptions) StreamsProducerOptions {
	return func(producer *StreamsProducer) {
		producer.producerOptions = poptions
	}
}

func (options StreamsProducerOptions) WithStreamOptions(soptions *stream.StreamOptions) StreamsProducerOptions {
	return func(producer *StreamsProducer) {
		producer.streamOptions = soptions
	}
}
