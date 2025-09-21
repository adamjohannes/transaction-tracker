package logger

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	logRepo "monthly-expenses-handler/internal/repository/log"
)

// New
// Returns a new slog logger configured to write to the database.
func New(pool *pgxpool.Pool) *slog.Logger {
	// Create the repository for logging
	repo := logRepo.NewPostgresRepository(pool)

	// Create our custom database handler
	handler := NewDBHandler(repo, slog.LevelInfo)

	return slog.New(handler)
}
