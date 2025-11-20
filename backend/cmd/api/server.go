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
func (app *dependencies) routes() http.Handler {
	mux := http.NewServeMux()

	// Auth routes (public)
	mux.HandleFunc("POST /register", app.registerUser)
	mux.HandleFunc("POST /login", app.loginUser)

	// Protected routes
	protectedMux := http.NewServeMux()
	protectedMux.HandleFunc("POST /transactions", app.createTransaction)
	protectedMux.HandleFunc("GET /transactions", app.listTransactions)

	// Public lookup routes
	lookupMux := http.NewServeMux()
	lookupMux.HandleFunc("GET /categories", app.listCategories)
	lookupMux.HandleFunc("GET /sub-categories", app.listAllSubCategories)
	lookupMux.HandleFunc("GET /categories/{category_name}/sub-categories", app.listSubCategoriesByCategory)
	lookupMux.HandleFunc("GET /status", app.listLookups(func() (any, error) {
		return app.statusController.GetAllStatus()
	}, "status"))
	lookupMux.HandleFunc("GET /currencies", app.listLookups(func() (any, error) {
		return app.currencyController.GetAllCurrencies()
	}, "currencies"))
	lookupMux.HandleFunc("GET /types", app.listLookups(func() (any, error) {
		return app.typeController.GetAllTransactionTypes()
	}, "types"))

	// Apply middleware
	var handler http.Handler = mux

	// Chain middlewares: Auth -> Logging
	protectedHandler := middleware.AuthMiddleware(protectedMux, app.authService)
	mux.Handle("/transactions", protectedHandler)
	mux.Handle("/transactions/", protectedHandler)

	// Mount public lookup routes
	mux.Handle("/categories", lookupMux)
	mux.Handle("/categories/", lookupMux)
	mux.Handle("/sub-categories", lookupMux)
	mux.Handle("/sub-categories/", lookupMux)
	mux.Handle("/status", lookupMux)
	mux.Handle("/currencies", lookupMux)
	mux.Handle("/types", lookupMux)

	// Apply logging middleware to all routes
	handler = middleware.LoggingMiddleware(mux, app.logger)

	return handler
}

// Serve
// Starts the HTTP server and handles graceful shutdown.
func (app *dependencies) Serve() {
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
