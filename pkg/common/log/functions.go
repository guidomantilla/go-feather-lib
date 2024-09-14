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

func retrieveSingleton() Logger {
	value := singleton.Load()
	if value == nil {
		return Default()
	}
	return value.(Logger)
}

func Default(writers ...io.Writer) Logger {
	logger := NewSlogLogger(SlogLevelOff.ValueFromName("INFO"), writers...)
	singleton.Store(logger)
	return logger
}

func Custom(writers ...io.Writer) Logger {
	logger := NewSlogLogger(SlogLevelOff.ValueFromName(strings.ToUpper(os.Getenv("LOG_LEVEL"))), writers...)
	singleton.Store(logger)
	return logger
}

//

func Debug(msg string, args ...any) {
	slogLogger := retrieveSingleton()
	slogLogger.Debug(context.Background(), msg, args...)
}

func Info(msg string, args ...any) {
	slogLogger := retrieveSingleton()
	slogLogger.Info(context.Background(), msg, args...)
}

func Warn(msg string, args ...any) {
	slogLogger := retrieveSingleton()
	slogLogger.Warn(context.Background(), msg, args...)
}

func Error(msg string, args ...any) {
	slogLogger := retrieveSingleton()
	slogLogger.Error(context.Background(), msg, args...)
}

func Fatal(msg string, args ...any) {
	slogLogger := retrieveSingleton()
	slogLogger.Fatal(context.Background(), msg, args...)
	os.Exit(1)
}

//

func AsSlogLogger() *slog.Logger {
	slogLogger := retrieveSingleton()
	return slogLogger.Logger().(*slog.Logger)
}
