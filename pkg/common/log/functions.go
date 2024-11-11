package log

import (
	"context"
	"log/slog"
	"os"
	"strings"
	"sync/atomic"
)

var singleton atomic.Value

func instance() Logger {
	value := singleton.Load()
	if value == nil {
		return Slog()
	}
	return value.(Logger)
}

func Slog(handlers ...slog.Handler) Logger {
	level := os.Getenv("LOG_LEVEL")
	if level == "" {
		level = "INFO"
	}
	logger := New(SlogLevelOff.ValueFromName(strings.ToUpper(level)), handlers...)
	singleton.Store(logger)
	return logger
}

//

func Trace(ctx context.Context, msg string, args ...any) {
	instance().Trace(ctx, msg, args...)
}

func Debug(ctx context.Context, msg string, args ...any) {
	instance().Debug(ctx, msg, args...)
}

func Info(ctx context.Context, msg string, args ...any) {
	instance().Info(ctx, msg, args...)
}

func Warn(ctx context.Context, msg string, args ...any) {
	instance().Warn(ctx, msg, args...)
}

func Error(ctx context.Context, msg string, args ...any) {
	instance().Error(ctx, msg, args...)
}

func Fatal(ctx context.Context, msg string, args ...any) {
	instance().Fatal(ctx, msg, args...)
	os.Exit(1)
}

//

func AsSlogLogger() *slog.Logger {
	return instance().Logger()
}
