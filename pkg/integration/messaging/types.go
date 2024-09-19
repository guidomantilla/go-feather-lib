package messaging

type MessageHandler interface {
	Handle(message Message[any], options ...ChannelOptions) error
}

type PollableChannel interface {
	Channel
	Receive(options ChannelOptions) (Message[any], error)
}

type SubscribableChannel interface {
	Channel
	Subscribe(handler MessageHandler, options ...ChannelOptions) error
	Unsubscribe(handler MessageHandler, options ...ChannelOptions) error
}
