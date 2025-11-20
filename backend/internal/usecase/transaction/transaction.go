package transaction

import (
	"context"
	"monthly-expenses-handler/internal/domain/transaction"
	"monthly-expenses-handler/internal/infrastructure/logger"
	transactionRepo "monthly-expenses-handler/internal/repository/transaction"
)

type UseCase struct {
	ctx             context.Context
	logger          *logger.Logger
	transactionRepo transactionRepo.Repository
}

// RecordTransaction
// Handles transaction insertion in the database.
func (u *UseCase) RecordTransaction(transaction *transaction.Transaction, userID uint64) (*transaction.Transaction, error) {
	newTransaction, err := u.transactionRepo.Create(u.ctx, transaction, userID)
	if err != nil {
		return nil, err
	}

	return newTransaction, nil
}

// List
// Handles transactions fetch.
func (u *UseCase) List(userID uint64) ([]*transaction.Transaction, error) {
	transactionsList, err := u.transactionRepo.List(u.ctx, userID)
	if err != nil {
		return nil, err
	}

	return transactionsList, nil
}

// Count
// Returns a map with the count of transactions.
func (u *UseCase) Count(groupBy string) (map[string]map[string]int, error) {
	countMap, err := u.transactionRepo.GetTransactionCount(u.ctx, groupBy)
	if err != nil {
		return nil, err
	}

	return countMap, nil
}

func (u *UseCase) CollectSubCategoryAmount() ([]transactionRepo.SubCategoryAmount, error) {
	amount, err := u.transactionRepo.GetSubCategoryAmounts(u.ctx)
	if err != nil {
		return nil, err
	}

	return amount, nil
}
