package logger

import (
	"log/slog"
	"os"
)

// New
// Returns a new slog logger configured to write JSON to standard output.
func New() *slog.Logger {
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}
	handler := slog.NewJSONHandler(os.Stdout, opts)
	return slog.New(handler)
}
