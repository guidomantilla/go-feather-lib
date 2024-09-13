package messaging

import amqp "github.com/rabbitmq/amqp091-go"

var _ RabbitMQQueueConnection = (*DefaultRabbitMQQueueConnection)(nil)

type RabbitMQQueueConnection interface {
	Start()
	Close()
	Connect() (*amqp.Connection, *amqp.Channel, *amqp.Queue, error)
}

var _ RabbitMQMessageListener = (*DefaultRabbitMQMessageListener)(nil)

type RabbitMQMessageListener interface {
	OnMessage(message *amqp.Delivery) error
}
