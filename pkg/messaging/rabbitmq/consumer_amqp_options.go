package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/guidomantilla/go-feather-lib/pkg/messaging"
)

var amqpConsumerOptions = NewAmqpConsumerOptions()

func NewAmqpConsumerOptions() AmqpConsumerOptions {
	return func(consumer messaging.Consumer) {
	}
}

type AmqpConsumerOptions func(consumer messaging.Consumer)

func (options AmqpConsumerOptions) WithRabbitMQListener(listener messaging.Listener[*amqp.Delivery]) AmqpConsumerOptions {
	return func(consumer messaging.Consumer) {
		consumer.Set("listener", listener)
	}
}

func (options AmqpConsumerOptions) WithAutoAck(autoAck bool) AmqpConsumerOptions {
	return func(consumer messaging.Consumer) {
		consumer.Set("autoAck", autoAck)
	}
}

func (options AmqpConsumerOptions) WithNoLocal(noLocal bool) AmqpConsumerOptions {
	return func(consumer messaging.Consumer) {
		consumer.Set("noLocal", noLocal)
	}
}

func (options AmqpConsumerOptions) WithDurable(durable bool) AmqpConsumerOptions {
	return func(consumer messaging.Consumer) {
		consumer.Set("durable", durable)
	}
}

func (options AmqpConsumerOptions) WithAutoDelete(autoDelete bool) AmqpConsumerOptions {
	return func(consumer messaging.Consumer) {
		consumer.Set("autoDelete", autoDelete)
	}
}

func (options AmqpConsumerOptions) WithExclusive(exclusive bool) AmqpConsumerOptions {
	return func(consumer messaging.Consumer) {
		consumer.Set("exclusive", exclusive)
	}
}

func (options AmqpConsumerOptions) WithNoWait(noWait bool) AmqpConsumerOptions {
	return func(consumer messaging.Consumer) {
		consumer.Set("noWait", noWait)
	}
}

func (options AmqpConsumerOptions) WithArgs(args amqp.Table) AmqpConsumerOptions {
	return func(consumer messaging.Consumer) {
		consumer.Set("args", args)
	}
}
