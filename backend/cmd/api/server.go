package api

import (
	"context"
	"log/slog"
	"monthly-expenses-handler/internal/middleware"
	"net/http"

	"monthly-expenses-handler/cmd/api/handlers"
	"monthly-expenses-handler/internal/controller/category"
	"monthly-expenses-handler/internal/controller/currency"
	"monthly-expenses-handler/internal/controller/status"
	"monthly-expenses-handler/internal/controller/sub_category"
	"monthly-expenses-handler/internal/controller/transaction"
	"monthly-expenses-handler/internal/controller/transaction_type"

	"github.com/jackc/pgx/v5/pgxpool"
)

// StartServer
// Initializes and runs the headless API server.
func StartServer(pool *pgxpool.Pool, logger *slog.Logger) {
	ctx := context.Background()
	mux := http.NewServeMux()

	// -- Initialize Controllers --
	txController := transaction.NewTransactionController(pool, ctx)
	categoryController := category.NewCategoryController(pool, ctx)
	subCategoryController := sub_category.NewSubCategoryController(pool, ctx)
	statusController := status.NewStatusController(pool, ctx)
	currencyController := currency.NewCurrencyController(pool, ctx)
	typeController := transaction_type.NewTypeController(pool, ctx)

	// -- Initialize Handlers --
	txHandler := handlers.NewTransactionHandler(txController, logger)
	categoryHandler := handlers.NewCategoryHandler(categoryController, logger)
	subCategoryHandler := handlers.NewSubCategoryHandler(subCategoryController, logger)
	statusHandler := handlers.NewStatusHandler(statusController)
	currencyHandler := handlers.NewCurrencyHandler(currencyController)
	typeHandler := handlers.NewTypeHandler(typeController)

	// -- Register Routes --
	// Transaction routes
	mux.HandleFunc("POST /transactions", txHandler.CreateTransaction)
	mux.HandleFunc("GET /transactions", txHandler.ListTransactions)

	// Category and Sub-Category routes
	mux.HandleFunc("GET /categories", categoryHandler.ListCategories)
	mux.HandleFunc("GET /sub-categories", subCategoryHandler.ListAllSubCategories)
	mux.HandleFunc("GET /categories/{category_name}/sub-categories", subCategoryHandler.ListSubCategoriesByCategory)

	// Lookup routes
	mux.HandleFunc("GET /status", statusHandler.ListStatus)
	mux.HandleFunc("GET /currencies", currencyHandler.ListCurrencies)
	mux.HandleFunc("GET /types", typeHandler.ListTypes)

	// Middleware
	var handler http.Handler = mux
	handler = middleware.LoggingMiddleware(handler, logger)

	// Start the HTTP server
	port := "8080"
	logger.Info("🚀 Starting API server", "port", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		logger.Error("Could not start server", "error", err)
	}
}
