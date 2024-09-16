package properties

import (
	"fmt"
	"strings"

	"github.com/guidomantilla/go-feather-lib/pkg/common/log"
)

type DefaultPropertiesOption func(properties *DefaultProperties)

func FromArray(array []string) DefaultPropertiesOption {
	return func(properties *DefaultProperties) {
		for _, env := range array {
			pair := strings.SplitN(env, "=", 2)
			if len(pair) != 2 {
				log.Error(fmt.Sprintf("[%s=??] not a key value parameter. expected [key=value]", pair[0]))
				continue
			}
			properties.Add(pair[0], pair[1])
		}
	}
}

type DefaultProperties struct {
	internalMap map[string]string
}

func NewDefaultProperties(options ...DefaultPropertiesOption) *DefaultProperties {
	properties := &DefaultProperties{
		internalMap: make(map[string]string),
	}

	for _, opt := range options {
		opt(properties)
	}

	return properties
}

func (p *DefaultProperties) Add(property string, value string) {
	if p.internalMap[property] == "" {
		p.internalMap[property] = value
	}
}

func (p *DefaultProperties) Get(property string) string {
	return p.internalMap[property]
}

func (p *DefaultProperties) AsMap() map[string]string {
	return p.internalMap
}
