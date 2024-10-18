package datasource

import (
	"strings"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type context_ struct {
	url     string
	server  string
	service string
}

func NewContext(url string, username string, password string, server string, service string) Context {

	if strings.TrimSpace(url) == "" {
		log.Fatal("starting up - error setting up datasource context: url is empty")
	}

	if strings.TrimSpace(username) == "" {
		log.Fatal("starting up - error setting up datasource context: username is empty")
	}

	if strings.TrimSpace(password) == "" {
		log.Fatal("starting up - error setting up datasource context: password is empty")
	}

	if strings.TrimSpace(server) == "" {
		log.Fatal("starting up - error setting up datasource context: server is empty")
	}

	if strings.TrimSpace(service) == "" {
		log.Fatal("starting up - error setting up datasource context: service is empty")
	}

	url = strings.Replace(url, ":username", username, 1)
	url = strings.Replace(url, ":password", password, 1)
	url = strings.Replace(url, ":server", server, 1)
	url = strings.Replace(url, ":service", service, 1)

	return &context_{
		url:     url,
		server:  server,
		service: service,
	}
}

func (context *context_) Url() string {
	return context.url
}

func (context *context_) Server() string {
	return context.server
}

func (context *context_) Service() string {
	return context.service
}
