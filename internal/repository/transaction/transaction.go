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
	Create(ctx context.Context, tx *transaction.Transaction) (int, error)
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
func (r *postgresRepository) Create(ctx context.Context, tx *transaction.Transaction) (int, error) {
	var id int
	query := `INSERT INTO transactions (amount, date, essential, status, currency, category, sub_category, description) 
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8) 
	          RETURNING id`

	// Use QueryRow for statements that are expected to return a single row.[5]
	err := r.db.QueryRow(ctx, query, tx.Amount.Abs(), tx.Essential, tx.Status.ID, tx.Currency.Code, tx.Category.Id, tx.SubCategory.ID, tx.Description).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to create transaction: %w", err)
	}

	return id, nil
}
