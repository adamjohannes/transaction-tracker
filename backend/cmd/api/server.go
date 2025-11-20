package api

import (
	"context"
	"monthly-expenses-handler/internal/dependency"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"monthly-expenses-handler/internal/middleware"

	"github.com/gin-gonic/gin"
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
	v1 := s.router.Group("/v1")

	{
		// Auth routes (public)
		auth := v1.Group("/auth")

		auth.POST("/register", s.deps.AuthController.Register)
		auth.POST("/login", s.deps.AuthController.Login)
	}

	{
		// Transaction routes (protected)
		transaction := v1.Group("/transaction")

		transaction.POST("/:userID", s.deps.TransactionController.PostTransaction)
		transaction.GET("/:userID", s.deps.TransactionController.GetAllTransactionsByUser)
	}

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
			s.deps.Logger.Error("Server error", map[string]interface{}{"error": err})
		}
	case sig := <-shutdownChan:
		s.deps.Logger.Info("Shutdown signal received", map[string]interface{}{"signal": sig})
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			s.deps.Logger.Error("Graceful shutdown failed", map[string]interface{}{"error": err})
		} else {
			s.deps.Logger.Info("Server shut down gracefully", nil)
		}
	}
}
