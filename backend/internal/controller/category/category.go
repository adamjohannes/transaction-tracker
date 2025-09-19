package category

import (
	"context"
	domain "monthly-expenses-handler/internal/domain/category"
	repository "monthly-expenses-handler/internal/repository/category"

	"github.com/jackc/pgx/v5/pgxpool"
)

// CategoryController
// Orchestrates category-related operations.
type CategoryController struct {
	pool *pgxpool.Pool
	ctx  context.Context
}

// NewCategoryController
// Creates a new instance of the category controller.
func NewCategoryController(pool *pgxpool.Pool, ctx context.Context) *CategoryController {
	return &CategoryController{
		pool: pool,
		ctx:  ctx,
	}
}

// GetAllCategories
// Fetches all categories from the repository.
func (cc *CategoryController) GetAllCategories() ([]*domain.Category, error) {
	repo := repository.NewPostgresRepository(cc.pool)
	categories, err := repo.GetAll(cc.ctx)

	if err != nil {
		return nil, err
	}

	return categories, nil
}
