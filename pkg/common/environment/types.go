package environment

import (
	"strconv"

	"github.com/guidomantilla/go-feather-lib/pkg/common/properties"
)

var (
	_ Environment = (*env)(nil)
	_ Environment = (*MockEnvironment)(nil)
)

type Environment interface {
	Value(property string) EnvVar
	ValueOrDefault(property string, defaultValue string) EnvVar
	PropertiesSources() []properties.PropertiesSource
	AppendPropertiesSources(propertySources ...properties.PropertiesSource)
}

func NewEnvironment(options ...EnvironmentOption) Environment {
	return newEnvironment(options...)
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
