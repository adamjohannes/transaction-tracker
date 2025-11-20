package sub_category

import (
	"context"
	"monthly-expenses-handler/internal/api_error"
	"monthly-expenses-handler/internal/infrastructure/logger"
	"monthly-expenses-handler/internal/usecase/sub_category"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Controller
// Orchestrates sub-category operations.
type Controller struct {
	ctx                context.Context
	subcategoryService sub_category.UseCase
	logger             *logger.Logger
}

// NewSubCategoryController
// Creates a new instance of the controller.
func NewSubCategoryController(ctx context.Context, subCategoryService sub_category.UseCase, logger *logger.Logger) *Controller {
	return &Controller{
		ctx,
		subCategoryService,
		logger,
	}
}

// GetAllSubCategories
// Fetches all sub-categories from the repository.
func (c *Controller) GetAllSubCategories(ctx *gin.Context) {
	c.logger.Info("Received a request to fetch all sub categories", nil)

	subCategoriesList, err := c.subcategoryService.List()
	if err != nil {
		c.logger.Error("Failed to fetch sub-categories", map[string]interface{}{"error": err})
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"subCategories": subCategoriesList,
	})
}

// GetSubCategoriesByCategory
// Fetches sub-categories filtered by the parent category name.
func (c *Controller) GetSubCategoriesByCategory(ctx *gin.Context) {
	c.logger.Info("Received a request to fetch sub categories by category", nil)

	categoryName := ctx.Param("category")
	if categoryName == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": api_error.NewValidationError("category name is required"),
		})
	}

	subCategoriesList, err := c.subcategoryService.ListByCategory(categoryName)
	if err != nil {
		c.logger.Error("Failed to fetch sub-categories by category", map[string]interface{}{"error": err})
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"subCategories": subCategoriesList,
	})
}
