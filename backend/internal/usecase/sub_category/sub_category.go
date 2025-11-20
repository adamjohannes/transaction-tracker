package sub_category

import (
	"context"
	domain "monthly-expenses-handler/internal/domain/sub_category"
	"monthly-expenses-handler/internal/infrastructure/logger"
	"monthly-expenses-handler/internal/repository/sub_category"
)

type UseCase struct {
	ctx          context.Context
	logger       *logger.Logger
	categoryRepo sub_category.Repository
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
