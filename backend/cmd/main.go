package main

import (
	"context"
	"monthly-expenses-handler/cmd/api"
	"monthly-expenses-handler/cmd/api/dependency"
	"monthly-expenses-handler/internal/config"
	"monthly-expenses-handler/internal/infrastructure/logger"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize context
	ctx := context.Background()

	// Initialize logger
	log := logger.New(cfg)

	// Build dependencies
	deps := dependencies.dependencies.BuildDependencies(cfg, ctx, log)

	// Build server
	server := api.NewServer(deps)

	// Start the server
	log.Info("Dependencies initialized, starting server...", nil)
	server.Serve()
}
