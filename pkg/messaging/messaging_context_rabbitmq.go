package messaging

import (
	"strings"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type RabbitMQContextOption func(rabbitMQContext *RabbitMQContext)

type RabbitMQContext struct {
	url    string
	server string
}

func NewRabbitMQContext(url string, username string, password string, server string, vhost string) *RabbitMQContext {

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

	if strings.TrimSpace(vhost) == "" {
		log.Fatal("starting up - error setting up rabbitMQContext: vhost is empty")
	}

	url = strings.Replace(url, ":username", username, 1)
	url = strings.Replace(url, ":password", password, 1)
	url = strings.Replace(url, ":server", server, 1)
	url = strings.Replace(url, ":vhost", vhost, 1)

	return &RabbitMQContext{
		url:    url,
		server: server + vhost,
	}
}

func (context *RabbitMQContext) Url() string {
	return context.url
}

func (context *RabbitMQContext) Server() string {
	return context.server
}
