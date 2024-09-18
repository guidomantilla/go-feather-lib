package datasource

import (
	"strings"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type OrmContext struct {
	url     string
	server  string
	service string
}

func NewOrmContext(url string, username string, password string, server string, service string) *OrmContext {

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

	return &OrmContext{
		url:     url,
		server:  server,
		service: service,
	}
}

func (context *OrmContext) Url() string {
	return context.url
}

func (context *OrmContext) Server() string {
	return context.server
}

func (context *OrmContext) Service() string {
	return context.service
}
