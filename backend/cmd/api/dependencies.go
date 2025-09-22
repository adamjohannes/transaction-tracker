package api

import (
	"log/slog"
	"monthly-expenses-handler/internal/controller/category"
	"monthly-expenses-handler/internal/controller/currency"
	"monthly-expenses-handler/internal/controller/status"
	"monthly-expenses-handler/internal/controller/sub_category"
	"monthly-expenses-handler/internal/controller/transaction"
	"monthly-expenses-handler/internal/controller/transaction_type"
)

// application
// Holds all the dependencies for the API.
type application struct {
	logger                *slog.Logger
	txController          *transaction.TransactionController
	categoryController    *category.CategoryController
	subCategoryController *sub_category.SubCategoryController
	statusController      *status.StatusController
	currencyController    *currency.CurrencyController
	typeController        *transaction_type.TypeController
}

// NewApplication
// Creates a new application instance with all dependencies.
func NewApplication(
	logger *slog.Logger,
	txController *transaction.TransactionController,
	categoryController *category.CategoryController,
	subCategoryController *sub_category.SubCategoryController,
	statusController *status.StatusController,
	currencyController *currency.CurrencyController,
	typeController *transaction_type.TypeController,
) *application {
	return &application{
		logger:                logger,
		txController:          txController,
		categoryController:    categoryController,
		subCategoryController: subCategoryController,
		statusController:      statusController,
		currencyController:    currencyController,
		typeController:        typeController,
	}
}
