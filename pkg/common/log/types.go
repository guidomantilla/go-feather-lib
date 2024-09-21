package log

import (
	"context"
	"io"
	"log/slog"
)

var (
	_ Logger[*slog.Logger] = (*slogLogger)(nil)
	_ Logger[any]          = (*MockLogger[any])(nil)
)

type Logger[T any] interface {
	Trace(ctx context.Context, msg string, args ...any)
	Debug(ctx context.Context, msg string, args ...any)
	Info(ctx context.Context, msg string, args ...any)
	Warn(ctx context.Context, msg string, args ...any)
	Error(ctx context.Context, msg string, args ...any)
	Fatal(ctx context.Context, msg string, args ...any)
	Logger() T
}

func NewLogger(level SlogLevel, writers ...io.Writer) Logger[*slog.Logger] {
	return newSlogLogger(level, writers...)
}
