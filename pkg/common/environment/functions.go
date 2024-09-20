package environment

import (
	"os"
	"sync/atomic"

	"github.com/guidomantilla/go-feather-lib/pkg/common/properties"
)

var singleton atomic.Value

func instance() Environment {
	value := singleton.Load()
	if value == nil {
		return Load()
	}
	return value.(Environment)
}

func Load(args ...[]string) Environment {
	withArgs := make([]EnvironmentOption, 0)
	withArgs = append(withArgs, WithSSL(), WithOs(os.Environ()))
	for _, arg := range args {
		withArgs = append(withArgs, WithCmd(arg))
	}
	env := NewEnvironment(withArgs...)
	singleton.Store(env)
	return env
}

func Value(property string) EnvVar {
	return instance().Value(property)
}

func ValueOrDefault(property string, defaultValue string) EnvVar {
	return instance().ValueOrDefault(property, defaultValue)
}

func PropertySources() []properties.PropertiesSource {
	return instance().PropertiesSources()
}

func AppendPropertySources(propertySources ...properties.PropertiesSource) {
	instance().AppendPropertiesSources(propertySources...)
}
