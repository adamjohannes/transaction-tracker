package currency

import (
	"context"
	"monthly-expenses-handler/internal/domain/currency"
	"monthly-expenses-handler/internal/infrastructure/logger"
	currencyRepo "monthly-expenses-handler/internal/repository/currency"
)

type UseCase struct {
	ctx          context.Context
	currencyRepo currencyRepo.Repository
	logger       *logger.Logger
}

func NewCurrencyService(ctx context.Context, currencyRepo currencyRepo.Repository, logger *logger.Logger) UseCase {
	return UseCase{
		ctx,
		currencyRepo,
		logger,
	}
}

func (u *UseCase) List() ([]*currency.Currency, error) {
	currencyList, err := u.currencyRepo.GetAll(u.ctx)

	if err != nil {
		return nil, err
	}

	return currencyList, nil
}
