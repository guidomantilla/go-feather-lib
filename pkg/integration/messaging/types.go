package messaging

import "time"

type Headers map[string]any

type Message interface {
	Headers() Headers
	Payload() any
}

type ChannelOptions func(channel Channel)

func WithTimeout(timeout time.Duration) ChannelOptions {
	return func(channel Channel) {
		channel.Timeout(timeout)
	}
}

//

type Channel interface {
	Send(message Message, options ...ChannelOptions) error
	Timeout(timeout time.Duration)
}

type MessageHandler interface {
	Handle(message Message, options ...ChannelOptions) error
}

type PollableChannel interface {
	Channel
	Receive(options ...ChannelOptions) (Message, error)
}

type SubscribableChannel interface {
	Channel
	Subscribe(handler MessageHandler, options ...ChannelOptions) error
	Unsubscribe(handler MessageHandler, options ...ChannelOptions) error
}
