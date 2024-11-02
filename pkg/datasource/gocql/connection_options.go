package gocql

import "github.com/gocql/gocql"

var connectionOptions = NewConnectionOptions()

func NewConnectionOptions() ConnectionOptions {
	return func(connection Connection) {
	}
}

type ConnectionOptions func(context Connection)

func (options ConnectionOptions) WithDialer(dialer gocql.HostDialer) ConnectionOptions {
	return func(connection Connection) {
		connection.Set("dialer", dialer)
	}
}
