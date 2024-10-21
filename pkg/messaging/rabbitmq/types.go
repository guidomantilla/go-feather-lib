package rabbitmq

import (
	"github.com/guidomantilla/go-feather-lib/pkg/messaging"
	amqp "github.com/rabbitmq/amqp091-go"
	samqp "github.com/rabbitmq/rabbitmq-stream-go-client/pkg/amqp"
	"github.com/rabbitmq/rabbitmq-stream-go-client/pkg/stream"
)

var (
	_ messaging.Context                         = (*context_)(nil)
	_ messaging.Connection[*amqp.Connection]    = (*Connection[*amqp.Connection])(nil)
	_ messaging.Connection[*stream.Environment] = (*Connection[*stream.Environment])(nil)
	_ messaging.Consumer                        = (*AmqpConsumer)(nil)
	_ messaging.Consumer                        = (*StreamsConsumer)(nil)
	_ messaging.Listener[*amqp.Delivery]        = (*AmqpListener)(nil)
	_ messaging.Listener[*samqp.Message]        = (*StreamsListener)(nil)
)
