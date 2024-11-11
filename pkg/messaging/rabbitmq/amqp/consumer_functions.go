package amqp

import (
	"context"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type ClosingHandler func(ctx context.Context, queue string, channel *amqp.Channel, deliveries <-chan amqp.Delivery, listener Listener, closeChannel chan string, fn MessageProcessor)

func closingHandler(ctx context.Context, queue string, channel *amqp.Channel, deliveries <-chan amqp.Delivery, listener Listener, closeChannel chan string, fn MessageProcessor) {
	var err error
	for message := range deliveries {
		go fn(ctx, listener, &message)
	}
	if err = channel.Close(); err != nil { //this line will be executed when the deliveries channel is closed
		log.Debug(ctx, fmt.Sprintf("rabbitmq consumer - failed to close channel to queue %s: %s", queue, err.Error()))
		return
	}
	close(closeChannel)
	log.Debug(ctx, fmt.Sprintf("rabbitmq consumer - disconected from queue %s", queue))
}

//

type MessageProcessor func(ctx context.Context, listener Listener, message *amqp.Delivery)

func messageProcessor(ctx context.Context, listener Listener, message *amqp.Delivery) {
	var err error
	log.Debug(ctx, fmt.Sprintf("rabbitmq consumer - message received: %s", message.Body))
	if err = listener.OnMessage(ctx, message); err != nil {
		log.Debug(ctx, fmt.Sprintf("rabbitmq consumer - failed to process message: %s", err.Error()))
		if err = message.Nack(false, true); err != nil {
			log.Debug(ctx, fmt.Sprintf("rabbitmq consumer - failed to nack message: %s", err.Error()))
		}
		log.Debug(ctx, fmt.Sprintf("rabbitmq consumer - nack message: %s", message.MessageId))
		return
	}
	if err = message.Ack(false); err != nil {
		log.Debug(ctx, fmt.Sprintf("rabbitmq consumer - failed to ack message: %s", err.Error()))
		return
	}
	log.Debug(ctx, fmt.Sprintf("rabbitmq consumer - ack message: %s", message.MessageId))
}
