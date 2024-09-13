package datasource

import (
	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
	"strings"
)

type DefaultDatasourceContext struct {
	url     string
	server  string
	service string
}

func NewDefaultDatasourceContext(url string, username string, password string, server string, service string) *DefaultDatasourceContext {

	if strings.TrimSpace(url) == "" {
		log.Fatal("starting up - error setting up datasourceContext: url is empty")
	}

	if strings.TrimSpace(username) == "" {
		log.Fatal("starting up - error setting up datasourceContext: username is empty")
	}

	if strings.TrimSpace(password) == "" {
		log.Fatal("starting up - error setting up datasourceContext: password is empty")
	}

	if strings.TrimSpace(server) == "" {
		log.Fatal("starting up - error setting up datasourceContext: server is empty")
	}

	if strings.TrimSpace(service) == "" {
		log.Fatal("starting up - error setting up datasourceContext: service is empty")
	}

	url = strings.Replace(url, ":username", username, 1)
	url = strings.Replace(url, ":password", password, 1)
	url = strings.Replace(url, ":server", server, 1)
	url = strings.Replace(url, ":service", service, 1)

	return &DefaultDatasourceContext{
		url:     url,
		server:  server,
		service: service,
	}
}

func (context *DefaultDatasourceContext) GetUrl() string {
	return context.url
}

func (context *DefaultDatasourceContext) GetServer() string {
	return context.server
}

func (context *DefaultDatasourceContext) GetService() string {
	return context.service
}
