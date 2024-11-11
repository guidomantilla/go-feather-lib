package log

import (
	"io"
	"log/slog"
)

const (
	SlogLevelTrace SlogLevel = iota
	SlogLevelDebug
	SlogLevelInfo
	SlogLevelWarning
	SlogLevelError
	SlogLevelFatal
	SlogLevelOff
)

type SlogLevel int

func (enum SlogLevel) String() string {

	switch enum {
	case SlogLevelTrace:
		return "TRACE"
	case SlogLevelDebug:
		return "DEBUG"
	case SlogLevelInfo:
		return "INFO"
	case SlogLevelWarning:
		return "WARN"
	case SlogLevelError:
		return "ERROR"
	case SlogLevelFatal:
		return "FATAL"
	case SlogLevelOff:
		return "OFF"
	}
	return "OFF"
}

func (enum SlogLevel) ToSlogLevel() slog.Level {
	switch enum {
	case SlogLevelTrace:
		return slog.Level(-8)
	case SlogLevelDebug:
		return slog.LevelDebug
	case SlogLevelInfo:
		return slog.LevelInfo
	case SlogLevelWarning:
		return slog.LevelWarn
	case SlogLevelError:
		return slog.LevelError
	case SlogLevelFatal:
		return slog.Level(12)
	case SlogLevelOff:
		return slog.Level(16)
	}
	return slog.Level(16)
}

func (enum SlogLevel) ValueFromName(slogLevel string) SlogLevel {
	switch slogLevel {
	case "TRACE":
		return SlogLevelTrace
	case "DEBUG":
		return SlogLevelDebug
	case "INFO":
		return SlogLevelInfo
	case "WARN":
		return SlogLevelWarning
	case "ERROR":
		return SlogLevelError
	case "FATAL":
		return SlogLevelFatal
	case "OFF":
		return SlogLevelOff
	}
	return SlogLevelOff
}

func (enum SlogLevel) ValueFromCardinal(slogLevel int) SlogLevel {
	switch slogLevel {
	case int(SlogLevelTrace):
		return SlogLevelTrace
	case int(SlogLevelDebug):
		return SlogLevelDebug
	case int(SlogLevelInfo):
		return SlogLevelInfo
	case int(SlogLevelWarning):
		return SlogLevelWarning
	case int(SlogLevelError):
		return SlogLevelError
	case int(SlogLevelFatal):
		return SlogLevelFatal
	case int(SlogLevelOff):
		return SlogLevelOff
	}
	return SlogLevelOff
}

func (enum SlogLevel) ValueFromSlogLevel(slogLevel slog.Level) SlogLevel {
	switch slogLevel {
	case slog.Level(-8):
		return SlogLevelTrace
	case slog.LevelDebug:
		return SlogLevelDebug
	case slog.LevelInfo:
		return SlogLevelInfo
	case slog.LevelWarn:
		return SlogLevelWarning
	case slog.LevelError:
		return SlogLevelError
	case slog.Level(12):
		return SlogLevelFatal
	case slog.Level(16):
		return SlogLevelOff
	}
	return SlogLevelOff
}

//

const (
	SlogTextFormat SlogFormat = iota
	SlogJsonFormat
)

type SlogFormat int

func (enum SlogFormat) String() string {
	switch enum {
	case SlogTextFormat:
		return "TEXT"
	case SlogJsonFormat:
		return "JSON"
	}
	return "TEXT"
}

func (enum SlogFormat) ValueFromName(loggerFormat string) SlogFormat {
	switch loggerFormat {
	case "TEXT":
		return SlogTextFormat
	case "JSON":
		return SlogJsonFormat
	}
	return SlogTextFormat
}

func (enum SlogFormat) ValueFromCardinal(loggerFormat int) SlogFormat {
	switch loggerFormat {
	case int(SlogTextFormat):
		return SlogTextFormat
	case int(SlogJsonFormat):
		return SlogJsonFormat
	}
	return SlogTextFormat
}

func (enum SlogFormat) Handler(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
	switch enum {
	case SlogTextFormat:
		return slog.NewTextHandler(w, opts)
	case SlogJsonFormat:
		return slog.NewJSONHandler(w, opts)
	}
	return slog.NewTextHandler(w, opts)
}
