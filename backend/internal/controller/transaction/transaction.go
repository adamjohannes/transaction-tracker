package transaction

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/net/context"
	"monthly-expenses-handler/internal/crypto"
	domain "monthly-expenses-handler/internal/domain/transaction"
	repository "monthly-expenses-handler/internal/repository/transaction"
)

// TransactionController
// Defines the interface for transaction data operations.
type TransactionController struct {
	pool   *pgxpool.Pool
	ctx    context.Context
	crypto *crypto.CryptoService
}

// NewTransactionController
// Creates a new instance of the transaction controller.
func NewTransactionController(pool *pgxpool.Pool, ctx context.Context, cryptoSvc *crypto.CryptoService) *TransactionController {
	return &TransactionController{
		pool:   pool,
		ctx:    ctx,
		crypto: cryptoSvc,
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

	// Initialize repository with database pool and crypto service
	repo := repository.NewPostgresRepository(tc.pool, tc.crypto)

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
	repo := repository.NewPostgresRepository(tc.pool, tc.crypto)
	return repo.GetAll(tc.ctx)
}

// GetFilteredTransactions
// Fetches transactions based on filter criteria.
func (tc *TransactionController) GetFilteredTransactions(filters *domain.FilterCriteria) ([]*domain.Transaction, error) {
	repo := repository.NewPostgresRepository(tc.pool, tc.crypto)
	return repo.GetFiltered(tc.ctx, filters)
}

// GetTransactionCount
// Fetches transaction counts grouped by a specific field.
func (tc *TransactionController) GetTransactionCount(groupBy string) (map[string]map[string]int, error) {
	repo := repository.NewPostgresRepository(tc.pool, tc.crypto)
	return repo.GetTransactionCount(tc.ctx, groupBy)
}

// GetSubCategoryAmounts
// Forwards the call to the repository.
func (tc *TransactionController) GetSubCategoryAmounts() ([]repository.SubCategoryAmount, error) {
	repo := repository.NewPostgresRepository(tc.pool, tc.crypto)
	return repo.GetSubCategoryAmounts(tc.ctx)
}
