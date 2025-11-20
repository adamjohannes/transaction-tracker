package category

import (
	"context"
	domain "monthly-expenses-handler/internal/domain/category"
	"monthly-expenses-handler/internal/infrastructure/logger"
	"monthly-expenses-handler/internal/repository/category"
)

type UseCase struct {
	ctx          context.Context
	logger       *logger.Logger
	categoryRepo category.Repository
}

func (u *UseCase) ListCategories() ([]*domain.Category, error) {
	categories, err := u.categoryRepo.GetAll(u.ctx)

	if err != nil {
		return nil, err
	}

	return categories, nil
}
