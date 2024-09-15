package environment

import (
	"strconv"

	"github.com/guidomantilla/go-feather-lib/pkg/common/properties"
)

var (
	_ Environment = (*DefaultEnvironment)(nil)
	_ Environment = (*MockEnvironment)(nil)
)

type Environment interface {
	Value(property string) EnvVar
	ValueOrDefault(property string, defaultValue string) EnvVar
	PropertySources() []properties.PropertySource
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
