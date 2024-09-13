package log

import (
	"context"
	"log/slog"
)

var (
	_ slog.Handler = (*FanoutHandler)(nil)
)

type FanoutHandler struct {
	handlers []slog.Handler
}

func NewFanoutHandler(handlers ...slog.Handler) slog.Handler {
	return &FanoutHandler{
		handlers: handlers,
	}
}

func (h *FanoutHandler) Enabled(ctx context.Context, l slog.Level) bool {
	for i := range h.handlers {
		if h.handlers[i].Enabled(ctx, l) {
			return true
		}
	}
	return false
}

func (h *FanoutHandler) Handle(ctx context.Context, r slog.Record) error {
	for i := range h.handlers {
		if h.handlers[i].Enabled(ctx, r.Level) {
			if err := h.handlers[i].Handle(ctx, r.Clone()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (h *FanoutHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	for _, handler := range h.handlers {
		handler.WithAttrs(attrs)
	}
	return NewFanoutHandler(h.handlers...)
}

func (h *FanoutHandler) WithGroup(name string) slog.Handler {
	for _, handler := range h.handlers {
		handler.WithGroup(name)
	}
	return NewFanoutHandler(h.handlers...)
}
