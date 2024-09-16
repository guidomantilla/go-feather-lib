package environment

import (
	"os"
	"sync/atomic"

	"github.com/guidomantilla/go-feather-lib/pkg/common/properties"
)

var singleton atomic.Value

func retrieveSingleton() Environment {
	value := singleton.Load()
	if value == nil {
		return Default()
	}
	return value.(Environment)
}

func Default() Environment {
	env := NewDefaultEnvironment(WithSSL(), WithOs(os.Environ()))
	singleton.Store(env)
	return env
}

func Custom(cmdArgs []string) Environment {
	env := NewDefaultEnvironment(WithSSL(), With(os.Environ(), cmdArgs))
	singleton.Store(env)
	return env
}

func Value(property string) EnvVar {
	env := retrieveSingleton()
	return env.Value(property)
}

func ValueOrDefault(property string, defaultValue string) EnvVar {
	env := retrieveSingleton()
	return env.ValueOrDefault(property, defaultValue)
}

func PropertySources() []properties.PropertySource {
	env := retrieveSingleton()
	return env.PropertySources()
}

func AppendPropertySources(propertySources ...properties.PropertySource) {
	env := retrieveSingleton()
	env.AppendPropertySources(propertySources...)
}
