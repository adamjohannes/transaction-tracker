package main

import (
	"context"
	"monthly-expenses-handler/cmd/api"
	"monthly-expenses-handler/internal/config"
	"monthly-expenses-handler/internal/logger"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize context
	ctx := context.Background()

	// Initialize logger
	logger := logger.New(cfg)

	// Build dependencies
	app := api.BuildDependencies(cfg, ctx, logger)

	// Start the server
	logger.Info("Dependencies initialized, starting server...", nil)
	app.Serve()
}
