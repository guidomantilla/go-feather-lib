package messaging

import (
	"time"

	"github.com/google/uuid"

	"github.com/guidomantilla/go-feather-lib/pkg/common/properties"
)

const (
	HeaderId           = "HeaderId"
	HeaderTimestamp    = "HeaderTimestamp"
	HeaderReplyChannel = "HeaderReplyChannel"
	HeaderErrorChannel = "HeaderErrorChannel"
)

type Headers interface {
	properties.Properties
	Id() uuid.UUID
	Timestamp() time.Time
	ReplyChannel() string
	ErrorChannel() string
}

type Message interface {
	Headers() Headers
	Payload() any
}

type ErrorMessage interface {
	Message
	Message()
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
	Timeout(timeout time.Duration) Channel
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
