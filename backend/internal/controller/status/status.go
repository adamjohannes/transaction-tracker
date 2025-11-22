package status

import (
	"context"
	"monthly-expenses-handler/internal/infrastructure/logger"
	statusService "monthly-expenses-handler/internal/usecase/status"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	ctx           context.Context
	statusService statusService.UseCase
	logger        *logger.Logger
}

func NewStatusController(ctx context.Context, statusService statusService.UseCase, logger *logger.Logger) *Controller {
	return &Controller{
		ctx,
		statusService,
		logger,
	}
}

func (c *Controller) GetAllStatus(ctx *gin.Context) {
	c.logger.Info("Received a request to fetch all status", nil)

	statusList, err := c.statusService.List()
	if err != nil {
		c.logger.Error("Failed to fetch status", map[string]interface{}{"error": err.Error()})
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": statusList,
	})
}
