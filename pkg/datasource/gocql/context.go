package gocql

import (
	"strings"

	"github.com/guidomantilla/go-feather-lib/pkg/common/assert"
)

type context_ struct {
	server []string
}

func NewContext(username string, password string, servers string) Context {
	assert.NotEmpty(username, "starting up - error setting up datasource context: username is empty")
	assert.NotEmpty(password, "starting up - error setting up datasource context: password is empty")
	assert.NotEmpty(servers, "starting up - error setting up datasource context: servers is empty")
	assert.True(strings.Contains(servers, ";"), "starting up - error setting up datasource context: servers is invalid")
	assert.True(len(strings.Split(servers, ";")) > 0, "starting up - error setting up datasource context: servers is empty")

	return &context_{
		server: strings.Split(servers, ";"),
	}
}

func (context *context_) Server() []string {
	return context.server
}
