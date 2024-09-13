package environment

import (
	"github.com/guidomantilla/go-feather-lib/pkg/common/properties"
	"strconv"
)

var _ Environment = (*DefaultEnvironment)(nil)

type Environment interface {
	GetValue(property string) EnvVar
	GetValueOrDefault(property string, defaultValue string) EnvVar
	GetPropertySources() []properties.PropertySource
	AppendPropertySources(propertySources ...properties.PropertySource)
}

//

type EnvVar string

func NewEnvVar(value string) EnvVar {
	return EnvVar(value)
}

func (envVar EnvVar) AsInt() (int, error) {
	value, err := strconv.Atoi(string(envVar))
	if err != nil {
		return 0, err
	}
	return value, nil
}

func (envVar EnvVar) AsString() string {
	return string(envVar)
}
