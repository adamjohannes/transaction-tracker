package database

import (
	"context"
	"log"
	"monthly-expenses-handler/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresDB struct {
	config *config.PostgresConfig
}

func New(config *config.PostgresConfig) *PostgresDB {
	return &PostgresDB{
		config: config,
	}
}

func (pdb *PostgresDB) Connect() (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), pdb.config.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to create connection pool: %v", err)
	}
	defer pool.Close()

	return pool, err
}
