package log

import (
	"context"
	"log/slog"
)

var (
	_ Logger = (*logger)(nil)
	_ Logger = (*MockLogger)(nil)
)

type Logger interface {
	Trace(ctx context.Context, msg string, args ...any)
	Debug(ctx context.Context, msg string, args ...any)
	Info(ctx context.Context, msg string, args ...any)
	Warn(ctx context.Context, msg string, args ...any)
	Error(ctx context.Context, msg string, args ...any)
	Fatal(ctx context.Context, msg string, args ...any)
	Handler() slog.Handler
	Logger() *slog.Logger
}
