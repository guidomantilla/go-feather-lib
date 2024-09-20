package environment

import (
	"os"
	"strings"

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

type EnvironmentOption func(environment Environment)

func WithCmd(cmdArgs []string) EnvironmentOption {
	return func(environment Environment) {
		cmdProperties := properties.NewProperties(properties.FromSlice(cmdArgs))
		environment.AppendPropertiesSources(properties.NewPropertiesSource(CmdPropertySourceName, cmdProperties))
	}
}

func WithSSL() EnvironmentOption {
	return func(environment Environment) {

		ValueOrEmpty := func(key string) string {
			if value, exists := os.LookupEnv(key); exists {
				return value
			}
			return ""
		}

		BuildOrEmpty := func(key string) string {
			if value := ValueOrEmpty(key); value != "" {
				return strings.Join([]string{os.Getenv("PWD"), "ssl", value}, "/")
			}
			return ""
		}

		sslProperties := properties.NewProperties()
		sslProperties.Add(SslServerName, ValueOrEmpty(SslServerName))
		sslProperties.Add(SslCaCertificate, BuildOrEmpty(SslCaCertificate))
		sslProperties.Add(SslClientCertificate, BuildOrEmpty(SslClientCertificate))
		sslProperties.Add(SslClientKey, BuildOrEmpty(SslClientKey))
		environment.AppendPropertiesSources(properties.NewPropertiesSource(SslPropertySourceName, sslProperties))
	}
}

func WithOs() EnvironmentOption {
	return func(environment Environment) {
		osProperties := properties.NewProperties(properties.FromSlice(os.Environ()))
		environment.AppendPropertiesSources(properties.NewPropertiesSource(OsPropertySourceName, osProperties))
	}
}

func With(cmdArgs []string) EnvironmentOption {
	return func(environment Environment) {
		WithCmd(cmdArgs)(environment)
		WithSSL()(environment)
		WithOs()(environment)
	}
}

func WithArraySource(name string, array []string) EnvironmentOption {
	return func(environment Environment) {
		environment.AppendPropertiesSources(properties.NewPropertiesSource(name, properties.NewProperties(properties.FromSlice(array))))
	}
}

func WithPropertySources(propertySources ...properties.PropertiesSource) EnvironmentOption {
	return func(environment Environment) {
		environment.AppendPropertiesSources(propertySources...)
	}
}

type env struct {
	environment
}

type environment struct {
	propertiesSources []properties.PropertiesSource
}

func newEnvironment(options ...EnvironmentOption) Environment {
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
