package database

import (
	"context"
	"fmt"
	"monthly-expenses-handler/internal/config"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

// ConnectDB establishes a connection pool to the PostgreSQL database.
// It builds the connection URI from individual environment variables.
func ConnectDB(cfg *config.PostgresConfig) (*pgxpool.Pool, error) {
	// Build the connection string from the config struct
	dbURL := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)

	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	return pool, nil
}

// getEnv is a helper function to read an environment variable with a fallback value.
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
