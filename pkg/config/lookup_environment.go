package config

import (
	envconfig "github.com/sethvargo/go-envconfig"

	"github.com/guidomantilla/go-feather-lib/pkg/environment"
)

var (
	_ envconfig.Lookuper = (*EnvironmentLookup)(nil)
)

type EnvironmentLookup struct {
	environment environment.Environment
}

func (lookuper *EnvironmentLookup) Lookup(key string) (string, bool) {
	value := lookuper.environment.GetValue(key).AsString()
	return value, value != ""
}
