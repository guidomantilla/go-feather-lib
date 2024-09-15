package messaging

import (
	"strings"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type RabbitMQContextOption func(rabbitMQContext *DefaultRabbitMQContext)

type DefaultRabbitMQContext struct {
	url    string
	server string
}

func NewDefaultRabbitMQContext(url string, username string, password string, server string) *DefaultRabbitMQContext {

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

	return &DefaultRabbitMQContext{
		url:    url,
		server: server,
	}
}

func (context *DefaultRabbitMQContext) Url() string {
	return context.url
}

func (context *DefaultRabbitMQContext) Server() string {
	return context.server
}
