package sub_category

import (
	"context"
	"fmt"
	domain "monthly-expenses-handler/internal/domain/sub_category"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Repository
// Defines the interface for sub-category data operations.
type Repository interface {
	GetAll(ctx context.Context) ([]*domain.SubCategory, error)
	GetByParentCategoryName(ctx context.Context, categoryName string) ([]*domain.SubCategory, error)
}

// postgresRepository
// Is the concrete implementation for PostgreSQL.
type postgresRepository struct {
	db *pgxpool.Pool
}

// NewPostgresRepository
// Creates a new instance of the sub-category repository.
func NewPostgresRepository(db *pgxpool.Pool) Repository {
	return &postgresRepository{db: db}
}

// GetAll
// Retrieves all sub-category records from the database.
func (r *postgresRepository) GetAll(ctx context.Context) ([]*domain.SubCategory, error) {
	query := `SELECT id, parent_category, name FROM transaction_sub_categories ORDER BY name`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query sub-categories: %w", err)
	}
	defer rows.Close()

	// Initialize as an empty slice
	subCategories := []*domain.SubCategory{}
	for rows.Next() {
		var id, parentID int8
		var name string
		if err := rows.Scan(&id, &parentID, &name); err != nil {
			return nil, fmt.Errorf("failed to scan sub-category row: %w", err)
		}
		subCategories = append(subCategories, domain.Build(id, parentID, name))
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("error reading sub-category rows: %w", rows.Err())
	}

	return subCategories, nil
}

// GetByParentCategoryName
// Retrieves sub-categories that belong to a given parent category.
func (r *postgresRepository) GetByParentCategoryName(ctx context.Context, categoryName string) ([]*domain.SubCategory, error) {
	query := `
		SELECT sc.id, sc.parent_category, sc.name
		FROM transaction_sub_categories sc
		JOIN transaction_categories tc ON sc.parent_category = tc.id
		WHERE tc.name = $1
		ORDER BY sc.name`

	rows, err := r.db.Query(ctx, query, categoryName)
	if err != nil {
		return nil, fmt.Errorf("failed to query sub-categories: %w", err)
	}
	defer rows.Close()

	// Initialize as an empty slice
	subCategories := []*domain.SubCategory{}
	for rows.Next() {
		var id, parentID int8
		var name string
		if err := rows.Scan(&id, &parentID, &name); err != nil {
			return nil, fmt.Errorf("failed to scan sub-category row: %w", err)
		}
		subCategories = append(subCategories, domain.Build(id, parentID, name))
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("error reading sub-category rows: %w", rows.Err())
	}

	return subCategories, nil
}
