package transaction_type

import (
	"context"
	"fmt"
	domain "monthly-expenses-handler/internal/domain/transaction_type"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	GetAll(ctx context.Context) ([]*domain.TransactionType, error)
}

type postgresRepository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(db *pgxpool.Pool) Repository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) GetAll(ctx context.Context) ([]*domain.TransactionType, error) {
	query := `SELECT id, name FROM transaction_types ORDER BY name`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query transaction types: %w", err)
	}
	defer rows.Close()

	// Initialize as an empty slice
	types := []*domain.TransactionType{}
	for rows.Next() {
		var id int8
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			return nil, fmt.Errorf("failed to scan transaction type row: %w", err)
		}
		types = append(types, domain.Build(id, name))
	}
	return types, rows.Err()
}
