package rabbitmq

import (
	"github.com/guidomantilla/go-feather-lib/pkg/messaging"
	amqp "github.com/rabbitmq/amqp091-go"
)

var amqpConsumerOptions = NewAmqpConsumerOptions()

func NewAmqpConsumerOptions() AmqpConsumerOptions {
	return func(consumer messaging.Consumer) {
	}
}

type AmqpConsumerOptions func(consumer messaging.Consumer)

func (option AmqpConsumerOptions) WithRabbitMQListener(listener messaging.Listener[*amqp.Delivery]) AmqpConsumerOptions {
	return func(consumer messaging.Consumer) {
		consumer.Set("listener", listener)
	}
}

func (option AmqpConsumerOptions) WithAutoAck(autoAck bool) AmqpConsumerOptions {
	return func(consumer messaging.Consumer) {
		consumer.Set("autoAck", autoAck)
	}
}

func (option AmqpConsumerOptions) WithNoLocal(noLocal bool) AmqpConsumerOptions {
	return func(consumer messaging.Consumer) {
		consumer.Set("noLocal", noLocal)
	}
}

func (option AmqpConsumerOptions) WithDurable(durable bool) AmqpConsumerOptions {
	return func(consumer messaging.Consumer) {
		consumer.Set("durable", durable)
	}
}

func (option AmqpConsumerOptions) WithAutoDelete(autoDelete bool) AmqpConsumerOptions {
	return func(consumer messaging.Consumer) {
		consumer.Set("autoDelete", autoDelete)
	}
}

func (option AmqpConsumerOptions) WithExclusive(exclusive bool) AmqpConsumerOptions {
	return func(consumer messaging.Consumer) {
		consumer.Set("exclusive", exclusive)
	}
}

func (option AmqpConsumerOptions) WithNoWait(noWait bool) AmqpConsumerOptions {
	return func(consumer messaging.Consumer) {
		consumer.Set("noWait", noWait)
	}
}

func (option AmqpConsumerOptions) WithArgs(args amqp.Table) AmqpConsumerOptions {
	return func(consumer messaging.Consumer) {
		consumer.Set("args", args)
	}
}
