package transaction

import (
	domain "monthly-expenses-handler/internal/domain/transaction"
	repository "monthly-expenses-handler/internal/repository/transaction"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/net/context"
)

// TransactionController
// Defines the interface for transaction data operations.
type TransactionController struct {
	pool *pgxpool.Pool
	ctx  context.Context
}

// NewTransactionController
// Creates a new instance of the transaction controller.
// It takes the database connection pool as a dependency.
func NewTransactionController(pool *pgxpool.Pool, ctx context.Context) *TransactionController {
	return &TransactionController{
		pool: pool,
		ctx:  ctx,
	}
}

// NewTransaction
// Inserts a new transaction record in the database.
// ERROR: Will fail if a required field is missing.
func (tc *TransactionController) NewTransaction(tempTransaction map[string]any) (*domain.Transaction, error) {
	// Build domain object from the raw map data
	transaction, err := domain.BuildTransaction(tempTransaction)
	if err != nil {
		return nil, err
	}

	// Initialize repository with database pool
	repo := repository.NewPostgresRepository(tc.pool)

	// Save the transaction.
	createdTransaction, err := repo.Create(tc.ctx, transaction)
	if err != nil {
		return nil, err
	}

	return createdTransaction, nil
}

// GetAllTransactions
// Fetches all transaction records from the database.
func (tc *TransactionController) GetAllTransactions() ([]*domain.Transaction, error) {
	// Initialize repository with database pool
	repo := repository.NewPostgresRepository(tc.pool)

	// Fetch all transactions.
	return repo.GetAll(tc.ctx)
}

// GetFilteredTransactions
// Fetches transactions based on filter criteria.
func (tc *TransactionController) GetFilteredTransactions(filters *domain.FilterCriteria) ([]*domain.Transaction, error) {
	repo := repository.NewPostgresRepository(tc.pool)
	return repo.GetFiltered(tc.ctx, filters)
}

// GetTransactionCount fetches transaction counts grouped by a specific field.
func (tc *TransactionController) GetTransactionCount(groupBy string) (map[string]map[string]int, error) {
	repo := repository.NewPostgresRepository(tc.pool)
	return repo.GetTransactionCount(tc.ctx, groupBy)
}
