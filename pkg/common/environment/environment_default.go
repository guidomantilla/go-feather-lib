package environment

import (
	properties "github.com/guidomantilla/go-feather-lib/pkg/common/properties"
)

const (
	OsPropertySourceName  = "OS_PROPERTY_SOURCE_NAME"
	CmdPropertySourceName = "CMD_PROPERTY_SOURCE_NAME" //nolint:gosec
)

type DefaultEnvironmentOption func(environment *DefaultEnvironment)

func WithOsArgsArray(osArgsArray []string) DefaultEnvironmentOption {
	return func(environment *DefaultEnvironment) {
		source := properties.NewDefaultPropertySource(OsPropertySourceName, properties.NewDefaultProperties(properties.FromArray(osArgsArray)))
		environment.propertySources = append(environment.propertySources, source)
	}
}

func WithCmdArgsArray(cmdArgsArray []string) DefaultEnvironmentOption {
	return func(environment *DefaultEnvironment) {
		source := properties.NewDefaultPropertySource(CmdPropertySourceName, properties.NewDefaultProperties(properties.FromArray(cmdArgsArray)))
		environment.propertySources = append(environment.propertySources, source)
	}
}

func WithArrays(osArgsArray []string, cmdArgsArray []string) DefaultEnvironmentOption {
	return func(environment *DefaultEnvironment) {
		osSource := properties.NewDefaultPropertySource(OsPropertySourceName, properties.NewDefaultProperties(properties.FromArray(osArgsArray)))
		cmdSource := properties.NewDefaultPropertySource(CmdPropertySourceName, properties.NewDefaultProperties(properties.FromArray(cmdArgsArray)))
		environment.propertySources = append(environment.propertySources, osSource, cmdSource)
	}
}

func WithArraySource(name string, array []string) DefaultEnvironmentOption {
	return func(environment *DefaultEnvironment) {
		source := properties.NewDefaultPropertySource(name, properties.NewDefaultProperties(properties.FromArray(array)))
		environment.propertySources = append(environment.propertySources, source)
	}
}

func WithPropertySources(propertySources ...properties.PropertySource) DefaultEnvironmentOption {
	return func(environment *DefaultEnvironment) {
		environment.propertySources = propertySources
	}
}

type DefaultEnvironment struct {
	propertySources []properties.PropertySource
}

func NewDefaultEnvironment(options ...DefaultEnvironmentOption) *DefaultEnvironment {
	environment := &DefaultEnvironment{
		propertySources: make([]properties.PropertySource, 0),
	}
	for _, opt := range options {
		opt(environment)
	}

	return environment
}

func (environment *DefaultEnvironment) GetValue(property string) EnvVar {

	var value string
	for _, source := range environment.propertySources {
		internalValue := source.Get(property)
		if internalValue != "" {
			value = internalValue
			break
		}
	}
	return NewEnvVar(value)
}

func (environment *DefaultEnvironment) GetValueOrDefault(property string, defaultValue string) EnvVar {

	envVar := environment.GetValue(property)
	if envVar != "" {
		return envVar
	}
	return NewEnvVar(defaultValue)
}

func (environment *DefaultEnvironment) GetPropertySources() []properties.PropertySource {
	return environment.propertySources
}

func (environment *DefaultEnvironment) AppendPropertySources(propertySources ...properties.PropertySource) {
	environment.propertySources = append(environment.propertySources, propertySources...)
}
