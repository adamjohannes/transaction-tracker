package transaction

import (
	"monthly-expenses-handler/internal/crypto"
	domain "monthly-expenses-handler/internal/domain/transaction"
	repository "monthly-expenses-handler/internal/repository/transaction"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/net/context"
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
func (tc *TransactionController) NewTransaction(tempTransaction map[string]any, userID int64) (*domain.Transaction, error) {
	transaction, err := domain.BuildTransaction(tempTransaction)
	if err != nil {
		return nil, err
	}

	repo := repository.NewPostgresRepository(tc.pool, tc.crypto)

	createdTransaction, err := repo.Create(tc.ctx, transaction, userID)
	if err != nil {
		return nil, err
	}

	return createdTransaction, nil
}

// GetAllTransactionsByUser
// Fetches all transaction records for a user.
func (tc *TransactionController) GetAllTransactionsByUser(userID int64) ([]*domain.Transaction, error) {
	repo := repository.NewPostgresRepository(tc.pool, tc.crypto)
	return repo.GetAllByUser(tc.ctx, userID)
}

// GetAllTransactions - DEPRECATED
// Fetches all transaction records from the database.
//func (tc *TransactionController) GetAllTransactions() ([]*domain.Transaction, error) {
//	repo := repository.NewPostgresRepository(tc.pool, tc.crypto)
//	return repo.GetAll(tc.ctx)
//}

// GetFilteredTransactions
// Fetches transactions based on filter criteria for a user.
func (tc *TransactionController) GetFilteredTransactions(filters *domain.FilterCriteria, userID int64) ([]*domain.Transaction, error) {
	repo := repository.NewPostgresRepository(tc.pool, tc.crypto)
	return repo.GetFiltered(tc.ctx, filters, userID)
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
