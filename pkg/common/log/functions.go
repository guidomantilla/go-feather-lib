package log

import (
	"context"
	"io"
	"log/slog"
	"os"
	"strings"
	"sync/atomic"
)

var singleton atomic.Value

func instance() Logger[*slog.Logger] {
	value := singleton.Load()
	if value == nil {
		return Slog()
	}
	return value.(Logger[*slog.Logger])
}

func Slog(writers ...io.Writer) Logger[*slog.Logger] {
	level := os.Getenv("LOG_LEVEL")
	if level == "" {
		level = "INFO"
	}
	logger := NewLogger(SlogLevelOff.ValueFromName(strings.ToUpper(level)), writers...)
	singleton.Store(logger)
	return logger
}

//

func Trace(msg string, args ...any) {
	instance().Trace(context.Background(), msg, args...)
}

func Debug(msg string, args ...any) {
	instance().Debug(context.Background(), msg, args...)
}

func Info(msg string, args ...any) {
	instance().Info(context.Background(), msg, args...)
}

func Warn(msg string, args ...any) {
	instance().Warn(context.Background(), msg, args...)
}

func Error(msg string, args ...any) {
	instance().Error(context.Background(), msg, args...)
}

func Fatal(msg string, args ...any) {
	instance().Fatal(context.Background(), msg, args...)
	os.Exit(1)
}

//

func AsSlogLogger() *slog.Logger {
	return instance().Logger()
}
