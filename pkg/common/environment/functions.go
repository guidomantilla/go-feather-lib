package environment

import (
	"github.com/guidomantilla/go-feather-lib/pkg/common/properties"
	"os"
	"sync/atomic"
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

func GetValue(property string) EnvVar {
	env := retrieveSingleton()
	return env.GetValue(property)
}

func GetValueOrDefault(property string, defaultValue string) EnvVar {
	env := retrieveSingleton()
	return env.GetValueOrDefault(property, defaultValue)
}

func GetPropertySources() []properties.PropertySource {
	env := retrieveSingleton()
	return env.GetPropertySources()
}

func AppendPropertySources(propertySources ...properties.PropertySource) {
	env := retrieveSingleton()
	env.AppendPropertySources(propertySources...)
}
