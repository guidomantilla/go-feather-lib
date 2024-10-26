package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

var amqpConsumerOptions = NewAmqpConsumerOptions()

func NewAmqpConsumerOptions() AmqpConsumerOptions {
	return func(consumer Consumer) {
	}
}

type AmqpConsumerOptions func(consumer Consumer)

func (options AmqpConsumerOptions) WithRabbitMQListener(listener Listener[*amqp.Delivery]) AmqpConsumerOptions {
	return func(consumer Consumer) {
		consumer.Set("listener", listener)
	}
}

func (options AmqpConsumerOptions) WithAutoAck(autoAck bool) AmqpConsumerOptions {
	return func(consumer Consumer) {
		consumer.Set("autoAck", autoAck)
	}
}

func (options AmqpConsumerOptions) WithNoLocal(noLocal bool) AmqpConsumerOptions {
	return func(consumer Consumer) {
		consumer.Set("noLocal", noLocal)
	}
}

func (options AmqpConsumerOptions) WithDurable(durable bool) AmqpConsumerOptions {
	return func(consumer Consumer) {
		consumer.Set("durable", durable)
	}
}

func (options AmqpConsumerOptions) WithAutoDelete(autoDelete bool) AmqpConsumerOptions {
	return func(consumer Consumer) {
		consumer.Set("autoDelete", autoDelete)
	}
}

func (options AmqpConsumerOptions) WithExclusive(exclusive bool) AmqpConsumerOptions {
	return func(consumer Consumer) {
		consumer.Set("exclusive", exclusive)
	}
}

func (options AmqpConsumerOptions) WithNoWait(noWait bool) AmqpConsumerOptions {
	return func(consumer Consumer) {
		consumer.Set("noWait", noWait)
	}
}

func (options AmqpConsumerOptions) WithArgs(args amqp.Table) AmqpConsumerOptions {
	return func(consumer Consumer) {
		consumer.Set("args", args)
	}
}
