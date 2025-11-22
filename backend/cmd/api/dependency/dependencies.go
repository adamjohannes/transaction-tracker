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
	"monthly-expenses-handler/internal/database"
	"monthly-expenses-handler/internal/infrastructure/crypto"
	"monthly-expenses-handler/internal/infrastructure/logger"
	categoryRepository "monthly-expenses-handler/internal/repository/category"
	currencyRepository "monthly-expenses-handler/internal/repository/currency"
	statusRepository "monthly-expenses-handler/internal/repository/status"
	subCategoryRepository "monthly-expenses-handler/internal/repository/sub_category"
	transactionRepo "monthly-expenses-handler/internal/repository/transaction"
	typeRepository "monthly-expenses-handler/internal/repository/transaction_type"
	userRepository "monthly-expenses-handler/internal/repository/user"
	authService "monthly-expenses-handler/internal/service/auth"
	categoryService "monthly-expenses-handler/internal/usecase/category"
	currencyService "monthly-expenses-handler/internal/usecase/currency"
	statusService "monthly-expenses-handler/internal/usecase/status"
	subCategoryService "monthly-expenses-handler/internal/usecase/sub_category"
	transactionService "monthly-expenses-handler/internal/usecase/transaction"
	transaction_type2 "monthly-expenses-handler/internal/usecase/transaction_type"
	userService "monthly-expenses-handler/internal/usecase/user"
)

type Dependencies struct {
	Logger *logger.Logger

	authService *authService.AuthService

	AuthController        *auth.Controller
	CategoryController    *category.Controller
	CurrencyController    *currency.Controller
	SubCategoryController *sub_category.Controller
	StatusController      *status.Controller
	TransactionController *transaction.Controller
	TypeController        *transaction_type.Controller
}

func BuildDependencies(cfg *config.Config, ctx context.Context, logger *logger.Logger) *Dependencies {
	// Connect to the database
	pool, err := database.ConnectDB(cfg, ctx, logger)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// Initialize Services
	cryptoSvc, err := crypto.NewCryptoService(cfg.Postgres.EncryptionKey)
	if err != nil {
		logger.Fatalf("Failed to create crypto service: %v", err)
	}

	authSvc := authService.NewAuthService(cfg.Postgres.JWTSecret, cfg.Postgres.SearchHashKey)
	userRepo := userRepository.NewPostgresRepository(pool, cryptoSvc, authSvc, logger)
	userSvc := userService.NewUserService(authSvc, ctx, logger, userRepo)

	categoryRepo := categoryRepository.NewPostgresRepository(pool)
	categorySvc := categoryService.NewCategoryService(ctx, categoryRepo, logger)

	currencyRepo := currencyRepository.NewPostgresRepository(pool)
	currencySvc := currencyService.NewCurrencyService(ctx, currencyRepo, logger)

	statusRepo := statusRepository.NewPostgresRepository(pool)
	statusSvc := statusService.NewStatusService(ctx, statusRepo, logger)

	subCategoryRepo := subCategoryRepository.NewPostgresRepository(pool)
	subCategorySvc := subCategoryService.NewSubCategoryService(ctx, subCategoryRepo, logger)

	transactionRepo := transactionRepo.NewPostgresRepository(pool, cryptoSvc)
	transactionSvc := transactionService.NewTransactionService(ctx, transactionRepo, logger)

	typeRepo := typeRepository.NewPostgresRepository(pool)
	typeSvc := transaction_type2.NewTypeService(ctx, typeRepo, logger)

	// Initialize Controllers
	authController := auth.NewAuthController(authSvc, userSvc, logger)
	categoryController := category.NewCategoryController(ctx, categorySvc, logger)
	currencyController := currency.NewCurrencyController(ctx, currencySvc, logger)
	statusController := status.NewStatusController(ctx, statusSvc, logger)
	subCategoryController := sub_category.NewSubCategoryController(ctx, subCategorySvc, logger)
	transactionController := transaction.NewTransactionController(transactionSvc, logger)
	typeController := transaction_type.NewTypeController(ctx, typeSvc, logger)

	return &Dependencies{
		logger,

		authSvc,

		authController,
		categoryController,
		currencyController,
		subCategoryController,
		statusController,
		transactionController,
		typeController,
	}
}
