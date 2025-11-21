package transaction

import (
	"errors"
	"monthly-expenses-handler/internal/api_error"
	"monthly-expenses-handler/internal/domain/transaction"
	"monthly-expenses-handler/internal/infrastructure/logger"
	"monthly-expenses-handler/internal/middleware"
	transactionService "monthly-expenses-handler/internal/usecase/transaction"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Controller
// Defines the interface for transaction data operations.
type Controller struct {
	transactionSvc transactionService.UseCase
	logger         *logger.Logger
}

type transactionPostRequest struct {
	Amount      string `json:"amount"`
	Category    string `json:"category"`
	Currency    string `json:"currency"`
	Date        string `json:"date"`
	Description string `json:"description"`
	Essential   bool   `json:"essential"`
	Status      string `json:"status"`
	SubCategory string `json:"subCategory"`
	Type        string `json:"type"`
}

// NewTransactionController
// Creates a new instance of the transaction controller.
func NewTransactionController(useCase transactionService.UseCase, logger *logger.Logger) *Controller {
	return &Controller{
		useCase,
		logger,
	}
}

// PostTransaction
// Handles the creation of a new transaction.
func (t *Controller) PostTransaction(c *gin.Context) {
	t.logger.Info("Received a request to record a new transaction", nil)

	userID, ok := c.Value(middleware.UserIDKey).(uint64)
	if !ok {
		t.logger.Error("user not authenticated", nil)
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "unauthorized",
			"detail": api_error.NewAuthError("user not authenticated"),
		})
		return
	}

	requestDatamap, err := collectPostRequest(c)
	if err != nil {
		t.logger.Error("Failed to bind request", map[string]interface{}{"error": err})
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed to bind request",
			"detail": err,
		})
		return
	}

	newTransaction, err := buildTransactionObj(requestDatamap)

	t.logger.Info("Attempting to create a new transaction...", nil)

	transactionObj, err := t.transactionSvc.RecordTransaction(newTransaction, userID)
	if err != nil {
		t.logger.Error("Failed to create transaction", map[string]interface{}{"error": err})

		var validationErr *api_error.ValidationError
		var authErr *api_error.ValidationError

		if errors.As(err, &validationErr) {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Failed to create transaction",
				"detail":  validationErr,
			})
		} else if errors.As(err, &authErr) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Failed to create transaction",
				"detail":  authErr,
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to create transaction",
				"detail":  err,
			})
		}

		return
	}

	t.logger.Info("Successfully created transaction", map[string]interface{}{"transaction_id": transactionObj.ID, "userID": userID})
	c.JSON(http.StatusCreated, gin.H{
		"message": "Successfully created transaction",
		"detail":  transactionObj,
	})
}

// GetAllTransactionsByUser
// Fetches all transaction records for a user.
func (t *Controller) GetAllTransactionsByUser(c *gin.Context) {
	t.logger.Info("Received a request to list transactions", nil)

	userID, ok := c.Value(middleware.UserIDKey).(uint64)
	if !ok {
		t.logger.Error("User not authenticated", nil)
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "user not authenticated",
			"detail":  api_error.NewAuthError("user not authenticated"),
		})
		return
	}

	t.logger.Info("Fetching transactions...", map[string]interface{}{"userID": userID})

	transactionList, err := t.transactionSvc.List(userID)
	if err != nil {
		t.logger.Error("Failed to list transactions", map[string]interface{}{"error": err})
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to list transactions",
			"detail":  err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"transactions": transactionList,
	})
}

// GetFilteredTransactions
// Fetches transactions based on filter criteria for a user.
func (t *Controller) GetFilteredTransactions(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "Not implemented",
	})
}

// GetTransactionCount
// Fetches transaction counts grouped by a specific field.
func (t *Controller) GetTransactionCount(c *gin.Context) {
	groupBy := c.Value(middleware.GroupBy).(string)

	countMap, err := t.transactionSvc.Count(groupBy)
	if err != nil {
		t.logger.Error("Failed to count transactions", map[string]interface{}{"error": err})
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to count transactions",
			"detail":  err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"count": countMap,
	})
}

// GetSubCategoryAmounts
// Forwards the call to the transactionRepo.
func (t *Controller) GetSubCategoryAmounts(c *gin.Context) {
	amount, err := t.transactionSvc.CollectSubCategoryAmount()
	if err != nil {
		t.logger.Error("Failed to collect sub-category amounts", map[string]interface{}{"error": err})
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to collect sub-category amounts",
			"detail":  err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"amount": amount,
	})
}

// --- Helpers

func collectPostRequest(c *gin.Context) (*transactionPostRequest, error) {
	var request *transactionPostRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func buildTransactionObj(request *transactionPostRequest) (*transaction.Transaction, error) {
	transactionMap := map[string]interface{}{
		"amount":      request.Amount,
		"category":    request.Category,
		"currency":    request.Currency,
		"date":        request.Date,
		"description": request.Description,
		"essential":   request.Essential,
		"status":      request.Status,
		"subCategory": request.SubCategory,
		"type":        request.Type,
	}

	return transaction.BuildTransaction(transactionMap)
}
