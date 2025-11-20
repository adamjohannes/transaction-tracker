package api

import (
	"context"
	"monthly-expenses-handler/internal/dependency"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

// setup
// Sets up the router for the API.
func (s *server) setup() {
	v1 := s.router.Group("/v1")

	// --- Protected routes
	{
		// Transaction routes
		transaction := v1.Group("/transaction")

		transaction.POST("/:userID", s.deps.TransactionController.PostTransaction)
		transaction.GET("/:userID", s.deps.TransactionController.GetAllTransactionsByUser)
	}

	// --- Public routes
	{
		// Auth routes (public)
		auth := v1.Group("/auth")

		auth.POST("/register", s.deps.AuthController.Register)
		auth.POST("/login", s.deps.AuthController.Login)
	}

	{
		// Lookup routes (public)
		lookup := v1.Group("/lookup")

		lookup.GET("/categories", s.deps.CategoryController.GetAllCategories)
		lookup.GET("/sub-categories", s.deps.SubCategoryController.GetAllSubCategories)
		lookup.GET("/sub-categories/:category_name", s.deps.SubCategoryController.GetSubCategoriesByCategory)
		lookup.GET("/status", s.deps.StatusController.GetAllStatus)
		lookup.HandleFunc("GET /currencies", s.deps.listLookups(func() (any, error) {
			return s.deps.CurrencyController.GetAllCurrencies()
		}, "currencies"))
		lookup.HandleFunc("GET /types", s.deps.listLookups(func() (any, error) {
			return s.deps.TypeController.GetAllTransactionTypes()
		}, "types"))
	}
}

// Serve
// Starts the HTTP server and handles graceful shutdown.
func (s *server) Serve() {
	port := "8080"
	server := &http.Server{
		Addr:    ":" + port,
		Handler: s.setup(),
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
