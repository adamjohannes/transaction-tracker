package status

import (
	"context"
	"fmt"
	domain "monthly-expenses-handler/internal/domain/status"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	GetAll(ctx context.Context) ([]*domain.Status, error)
}

type postgresRepository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(db *pgxpool.Pool) Repository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) GetAll(ctx context.Context) ([]*domain.Status, error) {
	query := `SELECT id, name FROM transaction_status ORDER BY name`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query status: %w", err)
	}
	defer rows.Close()

	var statuses []*domain.Status
	for rows.Next() {
		var id int8
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			return nil, fmt.Errorf("failed to scan status row: %w", err)
		}
		statuses = append(statuses, domain.Build(id, name))
	}
	return statuses, rows.Err()
}
