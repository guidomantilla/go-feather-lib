package messaging

import amqp "github.com/rabbitmq/amqp091-go"

var _ RabbitMQQueueConnection = (*DefaultRabbitMQQueueConnection)(nil)

type RabbitMQQueueConnection interface {
	Start()
	Close()
	Connect() (*amqp.Connection, *amqp.Channel, *amqp.Queue, error)
}

var _ RabbitMQQueueMessageListener = (*DefaultRabbitMQQueueMessageListener)(nil)

type RabbitMQQueueMessageListener interface {
	OnMessage(message *amqp.Delivery) error
}
