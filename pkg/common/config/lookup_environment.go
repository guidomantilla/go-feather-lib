package config

import (
	envconfig "github.com/sethvargo/go-envconfig"

	"github.com/guidomantilla/go-feather-lib/pkg/common/environment"
)

var (
	_ envconfig.Lookuper = (*EnvironmentLookup)(nil)
	_ envconfig.Lookuper = (*MockLookuper)(nil)
)

type EnvironmentLookup struct {
	environment environment.Environment
}

func (lookuper *EnvironmentLookup) Lookup(key string) (string, bool) {
	value := lookuper.environment.Value(key).AsString()
	return value, value != ""
}
