package config

import (
	"context"

	envconfig "github.com/sethvargo/go-envconfig"

	"github.com/guidomantilla/go-feather-lib/pkg/common/assert"
	"github.com/guidomantilla/go-feather-lib/pkg/common/environment"
)

func Process(ctx context.Context, environment environment.Environment, cfg *Config) error {
	assert.NotNil(ctx, "processing config - error processing config: context is nil")
	assert.NotNil(environment, "processing config - error processing config: environment is nil")
	assert.NotNil(cfg, "processing config - error processing config: config is nil")

	return envconfig.ProcessWith(ctx, &envconfig.Config{Target: cfg, Lookuper: &EnvironmentLookup{environment: environment}})
}
