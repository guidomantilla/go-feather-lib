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
	envs := os.Environ()
	env := NewDefaultEnvironment(WithArraySource(OsPropertySourceName, envs))
	singleton.Store(env)
	return env
}

func Custom(cmdArgsArray []string) Environment {
	envs := os.Environ()
	env := NewDefaultEnvironment(WithArrays(envs, cmdArgsArray))
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

func PsropertySources() []properties.PropertySource {
	env := retrieveSingleton()
	return env.PropertySources()
}

func AppendPropertySources(propertySources ...properties.PropertySource) {
	env := retrieveSingleton()
	env.AppendPropertySources(propertySources...)
}
