package log

import (
	"context"
)

var (
	_ Logger = (*SlogLogger)(nil)
	_ Logger = (*MockLogger)(nil)
)

type Logger interface {
	Debug(ctx context.Context, msg string, args ...any)
	Info(ctx context.Context, msg string, args ...any)
	Warn(ctx context.Context, msg string, args ...any)
	Error(ctx context.Context, msg string, args ...any)
	Fatal(ctx context.Context, msg string, args ...any)
	Logger() any
}
