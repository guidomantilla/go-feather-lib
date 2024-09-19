package messaging

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/guidomantilla/go-feather-lib/pkg/common/properties"
)

const (
	HeaderId           = "id"
	HeaderTimestamp    = "timestamp"
	HeaderReplyChannel = "reply-channel"
	HeaderErrorChannel = "error-channel"
)

type HeadersOptions func(headers Headers)

type HeadersOptionsChain interface {
	Build() HeadersOptions
	Id(id uuid.UUID) HeadersOptionsChain
	Timestamp(timestamp time.Time) HeadersOptionsChain
	ReplyChannel(replyChannel string) HeadersOptionsChain
	ErrorChannel(errorChannel string) HeadersOptionsChain
	Add(property string, value string) HeadersOptionsChain
}

type Headers interface {
	fmt.Stringer
	properties.Properties
	Id() uuid.UUID
	Timestamp() time.Time
	ReplyChannel() string
	ErrorChannel() string
}

type Message interface {
	fmt.Stringer
	Headers() Headers
	Payload() any
}

type ErrorMessage interface {
	Message
	Message()
}

//

type ChannelOptions func(channel Channel)

type ChannelOptionsChain interface {
	Build() ChannelOptions
	Timeout(timeout time.Duration) ChannelOptionsChain
}

type ChannelPipe chan<- Message

func (pipe ChannelPipe) String() string {
	return fmt.Sprintf("chan<- Message %d", cap(pipe))
}

type Channel interface {
	fmt.Stringer
	Send(message Message, options ...ChannelOptions) error
	Timeout(timeout time.Duration)
	Pipe(pipe ChannelPipe)
}

type MessageHandler interface {
	Handle(message Message, options ...ChannelOptions) error
}

type PollableChannel interface {
	Channel
	Receive(options ChannelOptions) (Message, error)
}

type SubscribableChannel interface {
	Channel
	Subscribe(handler MessageHandler, options ...ChannelOptions) error
	Unsubscribe(handler MessageHandler, options ...ChannelOptions) error
}
