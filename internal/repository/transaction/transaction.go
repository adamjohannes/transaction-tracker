package transaction

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

// TransactionRepository
// Handles database operations for transactions.
type TransactionRepository struct {
	pool *pgxpool.Pool
}

// New
// Creates a new transaction repository.
func New(pool *pgxpool.Pool) *TransactionRepository {
	return &TransactionRepository{
		pool: pool,
	}
}

// Create
// Inserts a new transaction into the database.
func (tr *TransactionRepository) Create() {
}
