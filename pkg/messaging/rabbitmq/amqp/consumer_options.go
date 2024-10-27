package amqp

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

var consumerOptions = NewConsumerOptions()

func NewConsumerOptions() ConsumerOptions {
	return func(consumer Consumer) {
	}
}

type ConsumerOptions func(consumer Consumer)

func (options ConsumerOptions) WithListener(listener Listener) ConsumerOptions {
	return func(consumer Consumer) {
		consumer.Set("listener", listener)
	}
}

func (options ConsumerOptions) WithAutoAck(autoAck bool) ConsumerOptions {
	return func(consumer Consumer) {
		consumer.Set("autoAck", autoAck)
	}
}

func (options ConsumerOptions) WithNoLocal(noLocal bool) ConsumerOptions {
	return func(consumer Consumer) {
		consumer.Set("noLocal", noLocal)
	}
}

func (options ConsumerOptions) WithDurable(durable bool) ConsumerOptions {
	return func(consumer Consumer) {
		consumer.Set("durable", durable)
	}
}

func (options ConsumerOptions) WithAutoDelete(autoDelete bool) ConsumerOptions {
	return func(consumer Consumer) {
		consumer.Set("autoDelete", autoDelete)
	}
}

func (options ConsumerOptions) WithExclusive(exclusive bool) ConsumerOptions {
	return func(consumer Consumer) {
		consumer.Set("exclusive", exclusive)
	}
}

func (options ConsumerOptions) WithNoWait(noWait bool) ConsumerOptions {
	return func(consumer Consumer) {
		consumer.Set("noWait", noWait)
	}
}

func (options ConsumerOptions) WithArgs(args amqp.Table) ConsumerOptions {
	return func(consumer Consumer) {
		consumer.Set("args", args)
	}
}

func (options ConsumerOptions) WithClosingHandler(closingHandler ClosingHandler) ConsumerOptions {
	return func(consumer Consumer) {
		consumer.Set("closingHandler", closingHandler)
	}
}

func (options ConsumerOptions) WithMessageProcessor(messageProcessor MessageProcessor) ConsumerOptions {
	return func(consumer Consumer) {
		consumer.Set("messageProcessor", messageProcessor)
	}
}
