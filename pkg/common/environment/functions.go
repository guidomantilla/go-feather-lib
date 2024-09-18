package environment

import (
	"os"
	"sync/atomic"

	"github.com/guidomantilla/go-feather-lib/pkg/common/properties"
)

var singleton atomic.Value

func retrieve() Environment {
	value := singleton.Load()
	if value == nil {
		return Load()
	}
	return value.(Environment)
}

func Load(args ...[]string) Environment {
	withArgs := make([]DefaultEnvironmentOption, 0)
	withArgs = append(withArgs, WithSSL(), WithOs(os.Environ()))
	for _, arg := range args {
		withArgs = append(withArgs, WithCmd(arg))
	}
	env := NewDefaultEnvironment(withArgs...)
	singleton.Store(env)
	return env
}

func Value(property string) EnvVar {
	env := retrieve()
	return env.Value(property)
}

func ValueOrDefault(property string, defaultValue string) EnvVar {
	env := retrieve()
	return env.ValueOrDefault(property, defaultValue)
}

func PropertySources() []properties.PropertySource {
	env := retrieve()
	return env.PropertySources()
}

func AppendPropertySources(propertySources ...properties.PropertySource) {
	env := retrieve()
	env.AppendPropertySources(propertySources...)
}
