package log

import (
	"context"
	rand "crypto/rand"
	"log/slog"
	"math/big"
	"slices"
	"time"

	"github.com/guidomantilla/go-feather-lib/pkg/common/errors"
)

var (
	_ slog.Handler = (*PoolHandler)(nil)
)

type PoolHandler struct {
	handlers []slog.Handler
}

func NewPoolHandler(handlers ...slog.Handler) slog.Handler {
	return &PoolHandler{
		handlers: handlers,
	}
}

func (h *PoolHandler) Enabled(ctx context.Context, l slog.Level) bool {
	for i := range h.handlers {
		if h.handlers[i].Enabled(ctx, l) {
			return true
		}
	}
	return false
}

func (h *PoolHandler) Handle(ctx context.Context, r slog.Record) error {
	var err error
	var random *big.Int
	if random, err = rand.Int(rand.Reader, big.NewInt(time.Now().UnixNano())); err != nil {
		return err
	}

	idx := random.Uint64() % uint64(len(h.handlers))
	handlers := append(h.handlers[idx:], h.handlers[:idx]...)

	var errs []error
	for i := range handlers {
		if handlers[i].Enabled(ctx, r.Level) {
			if err := h.handlers[i].Handle(ctx, r.Clone()); err != nil {
				errs = append(errs, err)
			}
		}
	}
	return errors.ErrJoin(errs...)
}

func (h *PoolHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	for _, handler := range h.handlers {
		handler.WithAttrs(slices.Clone(attrs))
	}
	return NewPoolHandler(h.handlers...)
}

func (h *PoolHandler) WithGroup(name string) slog.Handler {
	if name == "" {
		return h
	}
	for _, handler := range h.handlers {
		handler.WithGroup(name)
	}
	return NewPoolHandler(h.handlers...)
}
