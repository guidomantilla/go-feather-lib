package log

import (
	"context"
	"log/slog"
	"os"

	"github.com/guidomantilla/go-feather-lib/pkg/common/assert"
)

type ReplaceAttrFn func(groups []string, a slog.Attr) slog.Attr

func ReplaceAttr(_ []string, a slog.Attr) slog.Attr {
	if a.Key == slog.LevelKey {
		level := a.Value.Any().(slog.Level)
		a.Value = slog.StringValue(SlogLevelOff.ValueFromSlogLevel(level).String())
	}
	return a
}

//

type logger struct {
	internal *slog.Logger
}

func New(level SlogLevel, handlers ...slog.Handler) Logger {
	assert.NotNil(level, "starting up - error setting up logger: level is nil")

	if len(handlers) == 0 {
		opts := &slog.HandlerOptions{Level: level.ToSlogLevel(), ReplaceAttr: ReplaceAttr}
		handlers = append(handlers, SlogTextFormat.Handler(os.Stdout, opts))
	}

	internal := slog.New(NewFanoutHandler(handlers...))
	slog.SetDefault(internal)
	return &logger{internal: internal}
}

func (logger *logger) Trace(ctx context.Context, msg string, args ...any) {
	logger.internal.Log(ctx, SlogLevelTrace.ToSlogLevel(), msg, args...)
}

func (logger *logger) Debug(ctx context.Context, msg string, args ...any) {
	logger.internal.Log(ctx, SlogLevelDebug.ToSlogLevel(), msg, args...)
}

func (logger *logger) Info(ctx context.Context, msg string, args ...any) {
	logger.internal.Log(ctx, SlogLevelInfo.ToSlogLevel(), msg, args...)
}

func (logger *logger) Warn(ctx context.Context, msg string, args ...any) {
	logger.internal.Log(ctx, SlogLevelWarning.ToSlogLevel(), msg, args...)
}

func (logger *logger) Error(ctx context.Context, msg string, args ...any) {
	logger.internal.Log(ctx, SlogLevelError.ToSlogLevel(), msg, args...)
}

func (logger *logger) Fatal(ctx context.Context, msg string, args ...any) {
	logger.internal.Log(ctx, SlogLevelFatal.ToSlogLevel(), msg, args...)
	os.Exit(1)
}

func (logger *logger) Handler() slog.Handler {
	return logger.internal.Handler()
}

func (logger *logger) Logger() *slog.Logger {
	return logger.internal
}
