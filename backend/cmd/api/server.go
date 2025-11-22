package api

import (
	"context"
	"errors"
	"monthly-expenses-handler/cmd/api/dependency"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	deps       *dependencies.Dependencies
	router     *gin.Engine
	httpServer *http.Server
}

func NewServer(deps *dependencies.Dependencies) *Server {
	router := gin.New()
	srv := &http.Server{
		Handler: router,
	}

	return &Server{
		deps:       deps,
		router:     router,
		httpServer: srv,
	}
}

// setup
// Sets up the router for the API.
func (s *Server) setup() {
	v1 := s.router.Group("/v1")

	// --- Protected routes
	{
		// Transaction routes
		transaction := v1.Group("/transaction")

		transaction.GET("/:userID", s.deps.TransactionController.GetAllTransactionsByUser)
		transaction.POST("/:userID", s.deps.TransactionController.PostTransaction)
	}

	// --- Public routes
	{
		// Auth routes (public)
		auth := v1.Group("/auth")

		auth.POST("/login", s.deps.AuthController.Login)
		auth.POST("/register", s.deps.AuthController.Register)
	}

	{
		// Lookup routes (public)
		lookup := v1.Group("/lookup")

		lookup.GET("/categories", s.deps.CategoryController.GetAllCategories)
		lookup.GET("/currencies", s.deps.CurrencyController.GetAllCurrency)
		lookup.GET("/status", s.deps.StatusController.GetAllStatus)
		lookup.GET("/sub-categories", s.deps.SubCategoryController.GetAllSubCategories)
		lookup.GET("/sub-categories/:categoryName", s.deps.SubCategoryController.GetSubCategoriesByCategory)
		lookup.GET("/types", s.deps.TypeController.GetAllTransactionType)
	}
}

// Serve
// Starts the HTTP server and handles graceful shutdown.
func (s *Server) Serve() {
	port := "8080"
	s.httpServer.Addr = ":" + port
	s.setup()

	serverErrors := make(chan error, 1)
	go func() {
		s.deps.Logger.Info("Starting API server", map[string]interface{}{"port": port})
		if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErrors <- err
		}
	}()

	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		s.deps.Logger.Error("Server critical error", map[string]interface{}{"error": err.Error()})
	case sig := <-shutdownChan:
		s.deps.Logger.Info("Shutdown signal received", map[string]interface{}{"signal": sig.String()})
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := s.httpServer.Shutdown(shutdownCtx); err != nil {
			s.deps.Logger.Error("Graceful shutdown failed: Server forced to exit", map[string]interface{}{"error": err.Error()})
		} else {
			s.deps.Logger.Info("Server shut down gracefully", nil)
		}
	}
}
