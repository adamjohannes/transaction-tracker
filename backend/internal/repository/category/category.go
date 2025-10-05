package category

import (
	"context"
	"fmt"
	domain "monthly-expenses-handler/internal/domain/category"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Repository
// Defines the interface for category data operations.
type Repository interface {
	GetAll(ctx context.Context) ([]*domain.Category, error)
}

// postgresRepository
// Is the concrete implementation for PostgreSQL.
type postgresRepository struct {
	db *pgxpool.Pool
}

// NewPostgresRepository
// Creates a new instance of the category repository.
func NewPostgresRepository(db *pgxpool.Pool) Repository {
	return &postgresRepository{db: db}
}

// GetAll
// Retrieves all category records from the database.
func (r *postgresRepository) GetAll(ctx context.Context) ([]*domain.Category, error) {
	query := `SELECT id, name, description FROM transaction_categories ORDER BY name`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query categories: %w", err)
	}
	defer rows.Close()

	// Initialize as an empty slice
	categories := []*domain.Category{}
	for rows.Next() {
		var id int8
		var name, description string
		if err := rows.Scan(&id, &name, &description); err != nil {
			return nil, fmt.Errorf("failed to scan category row: %w", err)
		}
		// Use the domain's Build function to reconstruct the object
		categories = append(categories, domain.Build(id, name, description))
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("error reading category rows: %w", rows.Err())
	}

	return categories, nil
}
