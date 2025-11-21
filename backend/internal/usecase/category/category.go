package category

import (
	"context"
	domain "monthly-expenses-handler/internal/domain/category"
	"monthly-expenses-handler/internal/infrastructure/logger"
	"monthly-expenses-handler/internal/repository/category"
)

type UseCase struct {
	ctx          context.Context
	categoryRepo category.Repository
	logger       *logger.Logger
}

func NewCategoryService(ctx context.Context, categoryRepo category.Repository, logger *logger.Logger) *UseCase {
	return &UseCase{
		ctx,
		categoryRepo,
		logger,
	}
}

func (u *UseCase) ListCategories() ([]*domain.Category, error) {
	categories, err := u.categoryRepo.GetAll(u.ctx)

	if err != nil {
		return nil, err
	}

	return categories, nil
}
