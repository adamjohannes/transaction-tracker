package currency

import (
	"context"
	"monthly-expenses-handler/internal/infrastructure/logger"
	currencyService "monthly-expenses-handler/internal/usecase/currency"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	ctx             context.Context
	currencyService currencyService.UseCase
	logger          *logger.Logger
}

func NewCurrencyController(ctx context.Context, currencyService currencyService.UseCase, logger *logger.Logger) *Controller {
	return &Controller{
		ctx,
		currencyService,
		logger,
	}
}

func (c *Controller) GetAllCurrency(ctx *gin.Context) {
	c.logger.Info("Received a request to fetch all currency", nil)

	currencyList, err := c.currencyService.List()
	if err != nil {
		c.logger.Error("Failed to fetch currency", map[string]interface{}{"error": err})
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"currency": currencyList,
	})
}
