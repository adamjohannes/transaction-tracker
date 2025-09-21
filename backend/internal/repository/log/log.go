package log

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// LogEntry
// Defines the structure of a log entry.
type LogEntry struct {
	Level      string
	Message    string
	Attributes map[string]any
}

type Repository interface {
	Create(ctx context.Context, entry LogEntry) error
}

type postgresRepository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(db *pgxpool.Pool) Repository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) Create(ctx context.Context, entry LogEntry) error {
	// Marshal attributes map into a JSONB-compatible format
	attrBytes, err := json.Marshal(entry.Attributes)
	if err != nil {
		return fmt.Errorf("failed to marshal log attributes: %w", err)
	}

	query := `INSERT INTO app_logs (level, message, attributes) VALUES ($1, $2, $3)`
	_, err = r.db.Exec(ctx, query, entry.Level, entry.Message, attrBytes)

	if err != nil {
		return fmt.Errorf("failed to insert log entry: %w", err)
	}
	return nil
}
