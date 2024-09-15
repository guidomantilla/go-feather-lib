package messaging

import (
	"strings"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type NatsContextOption func(rabbitMQContext *NatsContext)

type NatsContext struct {
	url    string
	server string
}

// NewNatsContext creates a new NatsContext: nats://usuario:contraseña@localhost:4222
func NewNatsContext(url string, username string, password string, server string) *NatsContext {

	if strings.TrimSpace(url) == "" {
		log.Fatal("starting up - error setting up natsContext: url is empty")
	}

	if strings.TrimSpace(username) == "" {
		log.Fatal("starting up - error setting up natsContext: username is empty")
	}

	if strings.TrimSpace(password) == "" {
		log.Fatal("starting up - error setting up natsContext: password is empty")
	}

	if strings.TrimSpace(server) == "" {
		log.Fatal("starting up - error setting up natsContext: server is empty")
	}

	url = strings.Replace(url, ":username", username, 1)
	url = strings.Replace(url, ":password", password, 1)
	url = strings.Replace(url, ":server", server, 1)

	return &NatsContext{
		url:    url,
		server: server,
	}
}

func (context *NatsContext) Url() string {
	return context.url
}

func (context *NatsContext) Server() string {
	return context.server
}
