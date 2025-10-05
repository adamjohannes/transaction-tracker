package api

import (
	"log/slog"
	"monthly-expenses-handler/internal/auth"
	authCtrl "monthly-expenses-handler/internal/controller/auth"
	"monthly-expenses-handler/internal/controller/category"
	"monthly-expenses-handler/internal/controller/currency"
	"monthly-expenses-handler/internal/controller/status"
	"monthly-expenses-handler/internal/controller/sub_category"
	"monthly-expenses-handler/internal/controller/transaction"
	"monthly-expenses-handler/internal/controller/transaction_type"
	userCtrl "monthly-expenses-handler/internal/controller/user"
)

type application struct {
	logger                *slog.Logger
	txController          *transaction.TransactionController
	categoryController    *category.CategoryController
	subCategoryController *sub_category.SubCategoryController
	statusController      *status.StatusController
	currencyController    *currency.CurrencyController
	typeController        *transaction_type.TypeController
	authController        *authCtrl.AuthController
	userController        *userCtrl.UserController
	authService           *auth.AuthService
}

func NewApplication(
	logger *slog.Logger,
	txController *transaction.TransactionController,
	categoryController *category.CategoryController,
	subCategoryController *sub_category.SubCategoryController,
	statusController *status.StatusController,
	currencyController *currency.CurrencyController,
	typeController *transaction_type.TypeController,
	authController *authCtrl.AuthController,
	userController *userCtrl.UserController,
	authService *auth.AuthService,
) *application {
	return &application{
		logger:                logger,
		txController:          txController,
		categoryController:    categoryController,
		subCategoryController: subCategoryController,
		statusController:      statusController,
		currencyController:    currencyController,
		typeController:        typeController,
		authController:        authController,
		userController:        userController,
		authService:           authService,
	}
}
