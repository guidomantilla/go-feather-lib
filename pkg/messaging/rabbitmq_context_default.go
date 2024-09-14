package messaging

import (
	"strings"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type RabbitMQContextOption func(rabbitMQContext *DefaultRabbitMQContext)

func WithFailOver(failOver bool) RabbitMQContextOption {
	return func(rabbitMQContext *DefaultRabbitMQContext) {
		rabbitMQContext.failOver = failOver
	}
}

type DefaultRabbitMQContext struct {
	url                       string
	server                    string
	failOver                  bool
	notifyOnFaiOverConnection chan string
}

func NewDefaultRabbitMQContext(url string, username string, password string, server string, options ...RabbitMQContextOption) *DefaultRabbitMQContext {

	if strings.TrimSpace(url) == "" {
		log.Fatal("starting up - error setting up rabbitMQContext: url is empty")
	}

	if strings.TrimSpace(username) == "" {
		log.Fatal("starting up - error setting up rabbitMQContext: username is empty")
	}

	if strings.TrimSpace(password) == "" {
		log.Fatal("starting up - error setting up rabbitMQContext: password is empty")
	}

	if strings.TrimSpace(server) == "" {
		log.Fatal("starting up - error setting up rabbitMQContext: server is empty")
	}

	url = strings.Replace(url, ":username", username, 1)
	url = strings.Replace(url, ":password", password, 1)
	url = strings.Replace(url, ":server", server, 1)

	rabbitMQContext := &DefaultRabbitMQContext{
		url:                       url,
		server:                    server,
		notifyOnFaiOverConnection: make(chan string, 100),
	}

	for _, opt := range options {
		opt(rabbitMQContext)
	}

	return rabbitMQContext
}

func (context *DefaultRabbitMQContext) Url() string {
	return context.url
}

func (context *DefaultRabbitMQContext) Server() string {
	return context.server
}

func (context *DefaultRabbitMQContext) FailOver() bool {
	return context.failOver
}

func (context *DefaultRabbitMQContext) NotifyOnFaiOverConnection() chan string {
	return context.notifyOnFaiOverConnection
}
