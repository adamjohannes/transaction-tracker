package database

import (
	"context"
	"fmt"
	"monthly-expenses-handler/internal/config"
	"monthly-expenses-handler/internal/logger"

	"github.com/jackc/pgx/v5/pgxpool"
)

// ConnectDB establishes a connection pool to the PostgreSQL database.
// It builds the connection URI from individual environment variables.
func ConnectDB(cfg *config.Config, ctx context.Context, logger *logger.Logger) (*pgxpool.Pool, error) {
	// Build the connection string from config
	dbURL := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s",
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.Name,
	)

	logger.Debugf("Attempting connection with PostgreSQL at '%s'...", dbURL)
	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}

	logger.Debug("Successfully connected with PostgreSQL, pinging the DB...", nil)
	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	logger.Debug("Successfully pinged database", nil)
	return pool, nil
}
