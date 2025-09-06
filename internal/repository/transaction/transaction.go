package transaction

import (
	"context"
	"fmt"
	"monthly-expenses-handler/internal/domain/transaction"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Repository
// Defines the interface for transaction data operations.
type Repository interface {
	Create(ctx context.Context, tx *transaction.Transaction) (*transaction.Transaction, error)
}

// postgresRepository
// Is the concrete implementation of the Repository interface for PostgreSQL.
type postgresRepository struct {
	db *pgxpool.Pool
}

// NewPostgresRepository
// Creates a new instance of the transaction repository.
// It takes the database connection pool as a dependency.
func NewPostgresRepository(db *pgxpool.Pool) Repository {
	return &postgresRepository{db: db}
}

// Create
// Inserts a new transaction record into the database.
// It uses subqueries to look up foreign key IDs from names.
func (r *postgresRepository) Create(ctx context.Context, tx *transaction.Transaction) (*transaction.Transaction, error) {
	var id int
	query := `
		INSERT INTO transactions (amount, date, description, essential, type, status, currency, category, sub_category) 
		VALUES (
			$1, $2, $3, $4,
			(SELECT id FROM transaction_types WHERE name = $5),
			(SELECT id FROM transaction_status WHERE name = $6),
			(SELECT code FROM currencies WHERE code = $7),
			(SELECT id FROM transaction_categories WHERE name = $8),
			(SELECT id FROM transaction_sub_categories WHERE name = $9)
		) 
		RETURNING id`

	err := r.db.QueryRow(ctx, query,
		tx.Amount,
		tx.Date,
		tx.Description,
		tx.Essential,
		tx.Type.Name,
		tx.Status.Name,
		tx.Currency.Code,
		tx.Category.Name,
		tx.SubCategory.Name,
	).Scan(&id)

	if err != nil {
		return nil, fmt.Errorf("failed to create transaction: %w", err)
	}

	tx.ID = int8(id)
	return tx, nil
}
