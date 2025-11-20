package transaction_type

import (
	"context"
	"monthly-expenses-handler/internal/infrastructure/logger"
	transactionTypeService "monthly-expenses-handler/internal/usecase/transaction_type"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	ctx                    context.Context
	transactionTypeService transactionTypeService.UseCase
	logger                 *logger.Logger
}

func NewTypeController(ctx context.Context, transactionTypeService transactionTypeService.UseCase, logger *logger.Logger) *Controller {
	return &Controller{
		ctx,
		transactionTypeService,
		logger,
	}
}

func (c *Controller) GetAllTransactionType(ctx *gin.Context) {
	c.logger.Info("Received a request to fetch all transaction type", nil)

	transactionTypeList, err := c.transactionTypeService.List()
	if err != nil {
		c.logger.Error("Failed to fetch transaction type", map[string]interface{}{"error": err})
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"transactionType": transactionTypeList,
	})
}
