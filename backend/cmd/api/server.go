package api

import (
	"context"
	"monthly-expenses-handler/internal/dependencies"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"monthly-expenses-handler/internal/middleware"
)

type server struct {
	deps   *dependencies.Dependencies
	router *gin.Engine
}

func NewServer(deps *dependencies.Dependencies) *server {
	router := gin.New()

	return &server{
		deps,
		router,
	}
}

// routes
// Sets up the router for the API.
func (s *server) routes() http.Handler {
	s.router.Group("/v1")

	mux := http.NewServeMux()

	// Auth routes (public)
	mux.HandleFunc("POST /register", s.deps.registerUser)
	mux.HandleFunc("POST /login", s.deps.loginUser)

	// Protected routes
	protectedMux := http.NewServeMux()
	protectedMux.HandleFunc("POST /transactions", s.deps.createTransaction)
	protectedMux.HandleFunc("GET /transactions", s.deps.listTransactions)

	// Public lookup routes
	lookupMux := http.NewServeMux()
	lookupMux.HandleFunc("GET /categories", s.deps.listCategories)
	lookupMux.HandleFunc("GET /sub-categories", s.deps.listAllSubCategories)
	lookupMux.HandleFunc("GET /categories/{category_name}/sub-categories", s.deps.listSubCategoriesByCategory)
	lookupMux.HandleFunc("GET /status", s.deps.listLookups(func() (any, error) {
		return s.deps.StatusController.GetAllStatus()
	}, "status"))
	lookupMux.HandleFunc("GET /currencies", s.deps.listLookups(func() (any, error) {
		return s.deps.CurrencyController.GetAllCurrencies()
	}, "currencies"))
	lookupMux.HandleFunc("GET /types", s.deps.listLookups(func() (any, error) {
		return s.deps.TypeController.GetAllTransactionTypes()
	}, "types"))

	// Apply middleware
	var handler http.Handler = mux

	// Chain middlewares: Auth -> Logging
	protectedHandler := middleware.AuthMiddleware(protectedMux, s.deps.authService)
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
	handler = middleware.LoggingMiddleware(mux, s.deps.Logger)

	return handler
}

// Serve
// Starts the HTTP server and handles graceful shutdown.
func (s *server) Serve() {
	port := "8080"
	server := &http.Server{
		Addr:    ":" + port,
		Handler: s.routes(),
	}

	serverErrors := make(chan error, 1)
	go func() {
		s.deps.Logger.Info("🚀 Starting API server", map[string]interface{}{"port": port})
		serverErrors <- server.ListenAndServe()
	}()

	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		if err != nil && err != http.ErrServerClosed {
			s.deps.Logger.Error("Server error", "error", err)
		}
	case sig := <-shutdownChan:
		s.deps.Logger.Info("Shutdown signal received", "signal", sig)
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			s.deps.Logger.Error("Graceful shutdown failed", "error", err)
		} else {
			s.deps.Logger.Info("Server shut down gracefully")
		}
	}
}
