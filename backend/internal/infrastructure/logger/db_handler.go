package logger

import (
	"context"
	logRepo "monthly-expenses-handler/internal/repository/log"

	"log/slog"
)

// DBHandler
// A custom slog.Handler that writes log records to a database.
type DBHandler struct {
	repo   logRepo.Repository
	level  slog.Level
	attrs  []slog.Attr
	prefix string
}

func NewDBHandler(repo logRepo.Repository, level slog.Level) *DBHandler {
	return &DBHandler{
		repo:  repo,
		level: level,
	}
}

func (h *DBHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.level
}

func (h *DBHandler) Handle(ctx context.Context, r slog.Record) error {
	attributes := make(map[string]any)

	// Add attributes from the handler's context
	for _, attr := range h.attrs {
		attributes[attr.Key] = attr.Value.Any()
	}

	// Add attributes from the specific log call
	r.Attrs(func(attr slog.Attr) bool {
		attributes[attr.Key] = attr.Value.Any()
		return true
	})

	// Create the log entry and save it using the repository
	entry := logRepo.LogEntry{
		Level:      r.Level.String(),
		Message:    r.Message,
		Attributes: attributes,
	}

	// Use a background context if the request context is cancelled
	return h.repo.Create(context.Background(), entry)
}

// WithAttrs returns a new handler with the given attributes added.
func (h *DBHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	newHandler := *h
	newHandler.attrs = append(newHandler.attrs, attrs...)
	return &newHandler
}

// WithGroup returns a new handler with the given group name prepended.
func (h *DBHandler) WithGroup(name string) slog.Handler {
	newHandler := *h
	newHandler.prefix += name + "."
	return &newHandler
}
