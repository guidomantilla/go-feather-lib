package log

import (
	"context"
	"log/slog"
)

var (
	_ slog.Handler = (*SlogFanoutHandler)(nil)
)

type SlogFanoutHandler struct {
	handlers []slog.Handler
}

func NewSlogFanoutHandler(handlers ...slog.Handler) slog.Handler {
	return &SlogFanoutHandler{
		handlers: handlers,
	}
}

func (h *SlogFanoutHandler) Enabled(ctx context.Context, l slog.Level) bool {
	for i := range h.handlers {
		if h.handlers[i].Enabled(ctx, l) {
			return true
		}
	}
	return false
}

func (h *SlogFanoutHandler) Handle(ctx context.Context, r slog.Record) error {
	for i := range h.handlers {
		if h.handlers[i].Enabled(ctx, r.Level) {
			if err := h.handlers[i].Handle(ctx, r.Clone()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (h *SlogFanoutHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	for _, handler := range h.handlers {
		handler.WithAttrs(attrs)
	}
	return h
}

func (h *SlogFanoutHandler) WithGroup(name string) slog.Handler {
	for _, handler := range h.handlers {
		handler.WithGroup(name)
	}
	return h
}
