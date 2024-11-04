package goredis

var connectionOptions = NewConnectionOptions() //nolint:unused

func NewConnectionOptions() ConnectionOptions {
	return func(connection Connection) {
	}
}

type ConnectionOptions func(context Connection)
