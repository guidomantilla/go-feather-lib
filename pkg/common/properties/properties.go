package properties

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

func FromSlice(array []string) Option {
	return func(properties *properties) {
		for _, env := range array {
			pair := strings.SplitN(env, "=", 2)
			if len(pair) != 2 {
				log.Error(context.Background(), fmt.Sprintf("[%s=??] not a key value parameter. expected [key=value]", pair[0]))
				continue
			}
			properties.Add(pair[0], pair[1])
		}
	}
}

type properties struct {
	internal map[string]string
	mu       sync.RWMutex
}

func New(options ...Option) Properties {
	properties := &properties{
		internal: make(map[string]string),
	}

	for _, opt := range options {
		opt(properties)
	}

	return properties
}

func (p *properties) Add(property string, value string) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.internal[property] == "" {
		p.internal[property] = value
	}
}

func (p *properties) Get(property string) string {
	p.mu.Lock()
	defer p.mu.Unlock()

	return p.internal[property]
}

func (p *properties) AsMap() map[string]string {
	p.mu.Lock()
	defer p.mu.Unlock()

	return p.internal
}
