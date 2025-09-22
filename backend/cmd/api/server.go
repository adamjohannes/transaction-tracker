package api

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"monthly-expenses-handler/internal/middleware"
)

// routes
// Sets up the router for the API.
func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	// Transaction routes
	mux.HandleFunc("POST /transactions", app.createTransaction)
	mux.HandleFunc("GET /transactions", app.listTransactions)

	// Category and Sub-Category routes
	mux.HandleFunc("GET /categories", app.listCategories)
	mux.HandleFunc("GET /sub-categories", app.listAllSubCategories)
	mux.HandleFunc("GET /categories/{category_name}/sub-categories", app.listSubCategoriesByCategory)

	// Lookup routes
	mux.HandleFunc("GET /status", app.listLookups(func() (any, error) {
		return app.statusController.GetAllStatus()
	}, "status"))
	mux.HandleFunc("GET /currencies", app.listLookups(func() (any, error) {
		return app.currencyController.GetAllCurrencies()
	}, "currencies"))
	mux.HandleFunc("GET /types", app.listLookups(func() (any, error) {
		return app.typeController.GetAllTransactionTypes()
	}, "types"))

	// Apply middleware
	var handler http.Handler = mux
	handler = middleware.LoggingMiddleware(handler, app.logger)

	return handler
}

// Serve
// Starts the HTTP server and handles graceful shutdown.
func (app *application) Serve() {
	port := "8080"
	server := &http.Server{
		Addr:    ":" + port,
		Handler: app.routes(),
	}

	serverErrors := make(chan error, 1)
	go func() {
		app.logger.Info("🚀 Starting API server", "port", port)
		serverErrors <- server.ListenAndServe()
	}()

	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		if err != nil && err != http.ErrServerClosed {
			app.logger.Error("Server error", "error", err)
		}
	case sig := <-shutdownChan:
		app.logger.Info("Shutdown signal received", "signal", sig)
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			app.logger.Error("Graceful shutdown failed", "error", err)
		} else {
			app.logger.Info("Server shut down gracefully")
		}
	}
}
