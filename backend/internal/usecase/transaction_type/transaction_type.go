package transaction_type

import (
	"context"
	"monthly-expenses-handler/internal/domain/transaction_type"
	"monthly-expenses-handler/internal/infrastructure/logger"
	transactionTypeRepo "monthly-expenses-handler/internal/repository/transaction_type"
)

type UseCase struct {
	ctx                 context.Context
	transactionTypeRepo transactionTypeRepo.Repository
	logger              *logger.Logger
}

func NewTypeService(ctx context.Context, typeRepo transactionTypeRepo.Repository, logger *logger.Logger) *UseCase {
	return &UseCase{
		ctx,
		typeRepo,
		logger,
	}
}

func (u *UseCase) List() ([]*transaction_type.TransactionType, error) {
	transactionTypeList, err := u.transactionTypeRepo.GetAll(u.ctx)

	if err != nil {
		return nil, err
	}

	return transactionTypeList, nil
}
