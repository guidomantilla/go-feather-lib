package environment

import (
	properties "github.com/guidomantilla/go-feather-lib/pkg/common/properties"
)

const (
	SslPropertySourceName = "SSL_PROPERTY_SOURCE_NAME"
	OsPropertySourceName  = "OS_PROPERTY_SOURCE_NAME"
	CmdPropertySourceName = "CMD_PROPERTY_SOURCE_NAME" //nolint:gosec
)

const (
	SslServerName        = "SSL_SERVER_NAME"
	SslCaCertificate     = "SSL_CA_CERTIFICATE"
	SslClientCertificate = "SSL_CLIENT_CERTIFICATE"
	SslClientKey         = "SSL_CLIENT_KEY"
)

type environment struct {
	propertiesSources []properties.PropertiesSource
}

func New(options ...Options) Environment {
	environment := &environment{
		propertiesSources: make([]properties.PropertiesSource, 0),
	}
	for _, opt := range options {
		opt(environment)
	}

	return environment
}

func (environment *environment) Value(property string) EnvVar {

	var value string
	for _, source := range environment.propertiesSources {
		internalValue := source.Get(property)
		if internalValue != "" {
			value = internalValue
			break
		}
	}
	return NewEnvVar(value)
}

func (environment *environment) ValueOrDefault(property string, defaultValue string) EnvVar {

	envVar := environment.Value(property)
	if envVar != "" {
		return envVar
	}
	return NewEnvVar(defaultValue)
}

func (environment *environment) PropertiesSources() []properties.PropertiesSource {
	return environment.propertiesSources
}

func (environment *environment) AppendPropertiesSources(propertySources ...properties.PropertiesSource) {
	environment.propertiesSources = append(environment.propertiesSources, propertySources...)
}
