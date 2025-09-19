package currency

import (
	"context"
	"fmt"
	domain "monthly-expenses-handler/internal/domain/currency"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	GetAll(ctx context.Context) ([]*domain.Currency, error)
}

type postgresRepository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(db *pgxpool.Pool) Repository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) GetAll(ctx context.Context) ([]*domain.Currency, error) {
	query := `SELECT code FROM currencies ORDER BY code`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query currencies: %w", err)
	}
	defer rows.Close()

	var currencies []*domain.Currency
	for rows.Next() {
		var code string
		if err := rows.Scan(&code); err != nil {
			return nil, fmt.Errorf("failed to scan currency row: %w", err)
		}
		currency, _ := domain.New(code)
		currencies = append(currencies, currency)
	}
	return currencies, rows.Err()
}
