package sub_category

import (
	"context"
	domain "monthly-expenses-handler/internal/domain/sub_category"
	repository "monthly-expenses-handler/internal/repository/sub_category"

	"github.com/jackc/pgx/v5/pgxpool"
)

// SubCategoryController
// Orchestrates sub-category operations.
type SubCategoryController struct {
	pool *pgxpool.Pool
	ctx  context.Context
}

// NewSubCategoryController
// Creates a new instance of the controller.
func NewSubCategoryController(pool *pgxpool.Pool, ctx context.Context) *SubCategoryController {
	return &SubCategoryController{
		pool: pool,
		ctx:  ctx,
	}
}

// GetSubCategoriesByCategory
// Fetches sub-categories filtered by the parent category name.
func (scc *SubCategoryController) GetSubCategoriesByCategory(categoryName string) ([]*domain.SubCategory, error) {
	repo := repository.NewPostgresRepository(scc.pool)
	return repo.GetByParentCategoryName(scc.ctx, categoryName)
}

// GetAllSubCategories
// Fetches all sub-categories from the repository.
func (scc *SubCategoryController) GetAllSubCategories() ([]*domain.SubCategory, error) {
	repo := repository.NewPostgresRepository(scc.pool)
	return repo.GetAll(scc.ctx)
}
