package server

import "github.com/qmdx00/lifecycle"

type Application interface {
	ID() string
	Name() string
	Version() string
	Metadata() map[string]string
	Attach(name string, server lifecycle.Server)
	Run() error
}
