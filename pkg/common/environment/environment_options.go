package environment

import (
	"os"
	"strings"

	"github.com/guidomantilla/go-feather-lib/pkg/common/properties"
)

var options_ = NewOptions()

func NewOptions() Options {
	return func(environment Environment) {
	}
}

type Options func(environment Environment)

func (options Options) WithCmd(cmdArgs []string) Options {
	return func(environment Environment) {
		if len(cmdArgs) != 0 {
			cmdProperties := properties.New(properties.FromSlice(cmdArgs))
			environment.AppendPropertiesSources(properties.NewSource(CmdPropertySourceName, cmdProperties))
		}
	}
}

func (options Options) WithSSL() Options {
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

		sslProperties := properties.New()
		sslProperties.Add(SslServerName, ValueOrEmpty(SslServerName))
		sslProperties.Add(SslCaCertificate, BuildOrEmpty(SslCaCertificate))
		sslProperties.Add(SslClientCertificate, BuildOrEmpty(SslClientCertificate))
		sslProperties.Add(SslClientKey, BuildOrEmpty(SslClientKey))
		environment.AppendPropertiesSources(properties.NewSource(SslPropertySourceName, sslProperties))
	}
}

func (options Options) WithOs() Options {
	return func(environment Environment) {
		osProperties := properties.New(properties.FromSlice(os.Environ()))
		environment.AppendPropertiesSources(properties.NewSource(OsPropertySourceName, osProperties))
	}
}

func (options Options) WithArraySource(name string, array []string) Options {
	return func(environment Environment) {
		if strings.TrimSpace(name) != "" && len(array) != 0 {
			environment.AppendPropertiesSources(properties.NewSource(name, properties.New(properties.FromSlice(array))))

		}
	}
}

func (options Options) WithPropertySources(propertySources ...properties.PropertiesSource) Options {
	return func(environment Environment) {
		if len(propertySources) != 0 {
			environment.AppendPropertiesSources(propertySources...)
		}
	}
}
