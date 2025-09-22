package api

import (
	"context"
	"log/slog"
	"monthly-expenses-handler/internal/middleware"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	statusHandler := handlers.NewStatusHandler(statusController, logger)
	currencyHandler := handlers.NewCurrencyHandler(currencyController, logger)
	typeHandler := handlers.NewTypeHandler(typeController, logger)

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

	port := "8080"
	server := &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}

	// Channel to listen for errors from the server
	serverErrors := make(chan error, 1)
	go func() {
		logger.Info("🚀 Starting API server", "port", port)
		serverErrors <- server.ListenAndServe()
	}()

	// Channel to listen for interrupt signals from the OS
	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, syscall.SIGINT, syscall.SIGTERM)

	// Block until a signal or a server error is received
	select {
	case err := <-serverErrors:
		if err != nil && err != http.ErrServerClosed {
			logger.Error("Server error", "error", err)
		}
	case sig := <-shutdownChan:
		logger.Info("Shutdown signal received", "signal", sig)

		// Create a context with a 10-second timeout for shutdown
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Attempt to gracefully shut down the server
		if err := server.Shutdown(shutdownCtx); err != nil {
			logger.Error("Graceful shutdown failed", "error", err)
		} else {
			logger.Info("Server shut down gracefully")
		}
	}
}
