package log

import (
	"context"
	"io"
	"log/slog"
	"os"
	"sync/atomic"
)

var singleton atomic.Value

func retrieve() Logger {
	value := singleton.Load()
	if value == nil {
		return Slog()
	}
	return value.(Logger)
}

func Slog(writers ...io.Writer) Logger {
	level := os.Getenv("LOG_LEVEL")
	if level == "" {
		level = "INFO"
	}
	logger := NewSlogLogger(SlogLevelOff.ValueFromName(level), writers...)
	singleton.Store(logger)
	return logger
}

//

func Debug(msg string, args ...any) {
	slogLogger := retrieve()
	slogLogger.Debug(context.Background(), msg, args...)
}

func Info(msg string, args ...any) {
	slogLogger := retrieve()
	slogLogger.Info(context.Background(), msg, args...)
}

func Warn(msg string, args ...any) {
	slogLogger := retrieve()
	slogLogger.Warn(context.Background(), msg, args...)
}

func Error(msg string, args ...any) {
	slogLogger := retrieve()
	slogLogger.Error(context.Background(), msg, args...)
}

func Fatal(msg string, args ...any) {
	slogLogger := retrieve()
	slogLogger.Fatal(context.Background(), msg, args...)
	os.Exit(1)
}

//

func AsSlogLogger() *slog.Logger {
	slogLogger := retrieve()
	return slogLogger.Logger().(*slog.Logger)
}
