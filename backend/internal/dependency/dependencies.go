package dependencies

import (
	"context"
	"log"
	"monthly-expenses-handler/internal/config"
	"monthly-expenses-handler/internal/controller/auth"
	"monthly-expenses-handler/internal/controller/category"
	"monthly-expenses-handler/internal/controller/currency"
	"monthly-expenses-handler/internal/controller/status"
	"monthly-expenses-handler/internal/controller/sub_category"
	"monthly-expenses-handler/internal/controller/transaction"
	"monthly-expenses-handler/internal/controller/transaction_type"
	"monthly-expenses-handler/internal/controller/user"
	"monthly-expenses-handler/internal/database"
	"monthly-expenses-handler/internal/infrastructure/crypto"
	"monthly-expenses-handler/internal/infrastructure/logger"
	authSrvc "monthly-expenses-handler/internal/service/auth"
	userService2 "monthly-expenses-handler/internal/usecase/user"
)

type Dependencies struct {
	Logger *logger.Logger

	authService *authSrvc.AuthService

	AuthController        *auth.Controller
	CategoryController    *category.Controller
	CurrencyController    *currency.CurrencyController
	SubCategoryController *sub_category.Controller
	StatusController      *status.StatusController
	TransactionController *transaction.Controller
	TypeController        *transaction_type.TypeController
	UserController        *user.UserController
}

func BuildDependencies(cfg *config.Config, ctx context.Context, logger *logger.Logger) *Dependencies {
	// Connect to the database
	pool, err := database.ConnectDB(cfg, ctx, logger)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer pool.Close()

	// Initialize Services
	cryptoSvc, err := crypto.NewCryptoService(cfg.Postgres.EncryptionKey)
	if err != nil {
		logger.Fatalf("Failed to create crypto service: %v", err)
	}

	authService := authSrvc.NewAuthService(cfg.Postgres.JWTSecret, cfg.Postgres.SearchHashKey)
	userService := userService2.New(authService, ctx, logger, user)

	// Initialize Controllers
	userController := user.NewUserController(pool, cryptoSvc, authService)
	authController := auth.NewAuthController(authService, userController, ctx)
	transactionController := transaction.NewTransactionController(pool, ctx, cryptoSvc)
	categoryController := category.NewCategoryController(pool, ctx)
	subCategoryController := sub_category.NewSubCategoryController(pool, ctx)
	statusController := status.NewStatusController(pool, ctx)
	currencyController := currency.NewCurrencyController(pool, ctx)
	typeController := transaction_type.NewTypeController(pool, ctx)

	return &Dependencies{
		logger,

		authService,

		authController,
		categoryController,
		currencyController,
		subCategoryController,
		statusController,
		transactionController,
		typeController,
		userController,
	}
}
