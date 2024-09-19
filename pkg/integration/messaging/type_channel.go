package messaging

import (
	"fmt"
	"time"
)

type ChannelOptions func(channel Channel)

type ChannelOptionsChain interface {
	Build() ChannelOptions
	Timeout(timeout time.Duration) ChannelOptionsChain
}

type ChannelPipe chan<- Message[any]

func (pipe ChannelPipe) String() string {
	return fmt.Sprintf("chan<- Message %d", cap(pipe))
}

type Channel interface {
	fmt.Stringer
	Send(message Message[any], options ...ChannelOptions) error
	Timeout(timeout time.Duration)
	Pipe(pipe ChannelPipe)
}
