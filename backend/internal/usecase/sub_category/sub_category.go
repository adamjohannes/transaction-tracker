package sub_category

import (
	"context"
	domain "monthly-expenses-handler/internal/domain/sub_category"
	"monthly-expenses-handler/internal/infrastructure/logger"
	"monthly-expenses-handler/internal/repository/sub_category"
)

type UseCase struct {
	ctx          context.Context
	categoryRepo sub_category.Repository
	logger       *logger.Logger
}

func NewSubCategoryService(ctx context.Context, subCategoryRepo sub_category.Repository, logger *logger.Logger) *UseCase {
	return &UseCase{
		ctx,
		subCategoryRepo,
		logger,
	}
}

// List
// Handles fetching all sub-categories.
func (u *UseCase) List() ([]*domain.SubCategory, error) {
	subCategories, err := u.categoryRepo.GetAll(u.ctx)
	if err != nil {
		return nil, err
	}

	return subCategories, nil
}

// ListByCategory
// Handles fetching sub-categories for a parent.
func (u *UseCase) ListByCategory(category string) ([]*domain.SubCategory, error) {
	subCategories, err := u.categoryRepo.GetByParentCategoryName(u.ctx, category)
	if err != nil {
		return nil, err
	}

	return subCategories, nil
}
