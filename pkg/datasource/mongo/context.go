package mongo

import (
	"strings"

	"github.com/guidomantilla/go-feather-lib/pkg/common/assert"
)

type context_ struct {
	url    string
	server string
}

func NewContext(url string, username string, password string, server string) Context {
	assert.NotEmpty(url, "starting up - error setting up datasource context: url is empty")
	assert.NotEmpty(username, "starting up - error setting up datasource context: username is empty")
	assert.NotEmpty(password, "starting up - error setting up datasource context: password is empty")
	assert.NotEmpty(server, "starting up - error setting up datasource context: server is empty")

	url = strings.Replace(url, ":username", username, 1)
	url = strings.Replace(url, ":password", password, 1)
	url = strings.Replace(url, ":server", server, 1)

	return &context_{
		url:    url,
		server: server,
	}
}

func (context *context_) Url() string {
	return context.url
}

func (context *context_) Server() string {
	return context.server
}

/*
func (context *context_) Service() string {
	return context.service
}
*/
