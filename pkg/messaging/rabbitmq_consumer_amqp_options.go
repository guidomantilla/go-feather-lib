package messaging

import amqp "github.com/rabbitmq/amqp091-go"

var rabbitMQConsumerOptions = NewRabbitMQConsumerOptions()

func NewRabbitMQConsumerOptions() RabbitMQConsumerOptions {
	return func(consumer Consumer) {
	}
}

type RabbitMQConsumerOptions func(consumer Consumer)

func (option RabbitMQConsumerOptions) WithRabbitMQListener(listener Listener[*amqp.Delivery]) RabbitMQConsumerOptions {
	return func(consumer Consumer) {
		consumer.set("listener", listener)
	}
}

func (option RabbitMQConsumerOptions) WithAutoAck(autoAck bool) RabbitMQConsumerOptions {
	return func(consumer Consumer) {
		consumer.set("autoAck", autoAck)
	}
}

func (option RabbitMQConsumerOptions) WithNoLocal(noLocal bool) RabbitMQConsumerOptions {
	return func(consumer Consumer) {
		consumer.set("noLocal", noLocal)
	}
}

func (option RabbitMQConsumerOptions) WithDurable(durable bool) RabbitMQConsumerOptions {
	return func(consumer Consumer) {
		consumer.set("durable", durable)
	}
}

func (option RabbitMQConsumerOptions) WithAutoDelete(autoDelete bool) RabbitMQConsumerOptions {
	return func(consumer Consumer) {
		consumer.set("autoDelete", autoDelete)
	}
}

func (option RabbitMQConsumerOptions) WithExclusive(exclusive bool) RabbitMQConsumerOptions {
	return func(consumer Consumer) {
		consumer.set("exclusive", exclusive)
	}
}

func (option RabbitMQConsumerOptions) WithNoWait(noWait bool) RabbitMQConsumerOptions {
	return func(consumer Consumer) {
		consumer.set("noWait", noWait)
	}
}

func (option RabbitMQConsumerOptions) WithArgs(args amqp.Table) RabbitMQConsumerOptions {
	return func(consumer Consumer) {
		consumer.set("args", args)
	}
}
