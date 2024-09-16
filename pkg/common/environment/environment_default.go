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

type DefaultEnvironmentOption func(environment *DefaultEnvironment)

func WithCmd(cmdArgs []string) DefaultEnvironmentOption {
	return func(environment *DefaultEnvironment) {
		cmdProperties := properties.NewDefaultProperties(properties.FromSlice(cmdArgs))
		environment.propertySources = append(environment.propertySources, properties.NewDefaultPropertySource(CmdPropertySourceName, cmdProperties))
	}
}

func WithSSL() DefaultEnvironmentOption {
	return func(environment *DefaultEnvironment) {

		BuildOrEmpty := func(key string) string {
			if value, exists := os.LookupEnv(key); exists {
				return strings.Join([]string{os.Getenv("PWD"), "ssl", value}, "/")
			}
			return ""
		}

		sslProperties := properties.NewDefaultProperties()
		sslProperties.Add(SslServerName, BuildOrEmpty(os.Getenv(SslServerName)))
		sslProperties.Add(SslCaCertificate, BuildOrEmpty(os.Getenv(SslCaCertificate)))
		sslProperties.Add(SslClientCertificate, BuildOrEmpty(os.Getenv(SslClientCertificate)))
		sslProperties.Add(SslClientKey, BuildOrEmpty(os.Getenv(SslClientKey)))
		environment.propertySources = append(environment.propertySources, properties.NewDefaultPropertySource(SslPropertySourceName, sslProperties))
	}
}

func WithOs(osArgs []string) DefaultEnvironmentOption {
	return func(environment *DefaultEnvironment) {
		osProperties := properties.NewDefaultProperties(properties.FromSlice(osArgs))
		environment.propertySources = append(environment.propertySources, properties.NewDefaultPropertySource(OsPropertySourceName, osProperties))
	}
}

func With(osArgs []string, cmdArgs []string) DefaultEnvironmentOption {
	return func(environment *DefaultEnvironment) {
		WithCmd(cmdArgs)(environment)
		WithSSL()(environment)
		WithOs(osArgs)(environment)
	}
}

func WithArraySource(name string, array []string) DefaultEnvironmentOption {
	return func(environment *DefaultEnvironment) {
		source := properties.NewDefaultPropertySource(name, properties.NewDefaultProperties(properties.FromSlice(array)))
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

func (environment *DefaultEnvironment) Value(property string) EnvVar {

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

func (environment *DefaultEnvironment) ValueOrDefault(property string, defaultValue string) EnvVar {

	envVar := environment.Value(property)
	if envVar != "" {
		return envVar
	}
	return NewEnvVar(defaultValue)
}

func (environment *DefaultEnvironment) PropertySources() []properties.PropertySource {
	return environment.propertySources
}

func (environment *DefaultEnvironment) AppendPropertySources(propertySources ...properties.PropertySource) {
	environment.propertySources = append(environment.propertySources, propertySources...)
}
