package gorm

import (
	"strings"

	"github.com/guidomantilla/go-feather-lib/pkg/common/assert"
)

type context_ struct {
	url      string
	username string
	password string
	server   string
	service  string
}

func NewContext(url string, username string, password string, server string, service string) Context {
	assert.NotEmpty(url, "starting up - error setting up datasource context: url is empty")
	assert.NotEmpty(username, "starting up - error setting up datasource context: username is empty")
	assert.NotEmpty(password, "starting up - error setting up datasource context: password is empty")
	assert.NotEmpty(server, "starting up - error setting up datasource context: server is empty")
	assert.NotEmpty(service, "starting up - error setting up datasource context: service is empty")

	url = strings.Replace(url, ":username", username, 1)
	url = strings.Replace(url, ":password", password, 1)
	url = strings.Replace(url, ":server", server, 1)
	url = strings.Replace(url, ":service", service, 1)

	return &context_{
		url:      url,
		username: username,
		password: password,
		server:   server,
		service:  service,
	}
}

func (context *context_) Url() string {
	return context.url
}

func (context *context_) User() string {
	return context.username
}

func (context *context_) Password() string {
	return context.password
}

func (context *context_) Server() any {
	return context.server
}

func (context *context_) Service() string {
	return context.service
}
