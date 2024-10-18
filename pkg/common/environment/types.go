package environment

import (
	"strconv"

	"github.com/guidomantilla/go-feather-lib/pkg/common/properties"
)

var (
	_ Environment = (*environment)(nil)
	_ Environment = (*MockEnvironment)(nil)
)

type Environment interface {
	Value(property string) EnvVar
	ValueOrDefault(property string, defaultValue string) EnvVar
	PropertiesSources() []properties.PropertiesSource
	AppendPropertiesSources(propertySources ...properties.PropertiesSource)
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
