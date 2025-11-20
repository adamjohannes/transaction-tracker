package category

import (
	"context"
	"monthly-expenses-handler/internal/infrastructure/logger"
	"monthly-expenses-handler/internal/usecase/category"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Controller
// Orchestrates category-related operations.
type Controller struct {
	ctx             context.Context
	categoryService category.UseCase
	logger          *logger.Logger
}

// NewCategoryController
// Creates a new instance of the category controller.
func NewCategoryController(ctx context.Context, categoryService category.UseCase, logger *logger.Logger) *Controller {
	return &Controller{
		ctx,
		categoryService,
		logger,
	}
}

// GetAllCategories
// Fetches all categories from the repository.
func (cc *Controller) GetAllCategories(ctx *gin.Context) {
	cc.logger.Info("Received a request to fetch all categories", nil)

	categories, err := cc.categoryService.ListCategories()
	if err != nil {
		cc.logger.Error("Failed to fetch categories", map[string]interface{}{"error": err})
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to fetch categories",
			"detail":  err,
		})
		return
	}

	cc.logger.Info("Successfully fetched all categories", map[string]interface{}{"count": len(categories)})
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message":    "Successfully fetched all categories",
		"categories": categories,
	})
}
