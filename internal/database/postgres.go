package database

import (
	"context"
	"fmt"
	"monthly-expenses-handler/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Connect establishes a new connection pool to the database.
// It returns the pool or an error if the connection fails.
func Connect(cfg *config.PostgresConfig) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), cfg.DatabaseURL)

	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	return pool, nil
}
