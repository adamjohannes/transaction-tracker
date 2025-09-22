package main

import (
	"context"
	"log"
	"monthly-expenses-handler/cmd/api"
	"monthly-expenses-handler/internal/config"
	"monthly-expenses-handler/internal/controller/category"
	"monthly-expenses-handler/internal/controller/currency"
	"monthly-expenses-handler/internal/controller/status"
	"monthly-expenses-handler/internal/controller/sub_category"
	"monthly-expenses-handler/internal/controller/transaction"
	"monthly-expenses-handler/internal/controller/transaction_type"
	"monthly-expenses-handler/internal/database"
	"monthly-expenses-handler/internal/logger"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Could not load configuration: %v", err)
	}

	// Connect to the database
	pool, err := database.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer pool.Close()

	// Initialize dependencies
	slogLogger := logger.New(pool)
	ctx := context.Background()

	// -- Initialize Controllers --
	txController := transaction.NewTransactionController(pool, ctx)
	categoryController := category.NewCategoryController(pool, ctx)
	subCategoryController := sub_category.NewSubCategoryController(pool, ctx)
	statusController := status.NewStatusController(pool, ctx)
	currencyController := currency.NewCurrencyController(pool, ctx)
	typeController := transaction_type.NewTypeController(pool, ctx)

	// -- Create the application instance --
	app := api.NewApplication(
		slogLogger,
		txController,
		categoryController,
		subCategoryController,
		statusController,
		currencyController,
		typeController,
	)

	// -- Start the server --
	slogLogger.Info("Dependencies initialized, starting server")
	app.Serve()
}
