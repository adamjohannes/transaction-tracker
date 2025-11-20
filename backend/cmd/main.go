package main

import (
	"context"
	"monthly-expenses-handler/cmd/api"
	"monthly-expenses-handler/internal/config"
	"monthly-expenses-handler/internal/dependencies"
	"monthly-expenses-handler/internal/logger"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize context
	ctx := context.Background()

	// Initialize logger
	log := logger.New(cfg)

	// Build dependencies
	deps := dependencies.BuildDependencies(cfg, ctx, log)

	// Build server
	server := api.NewServer(deps)

	// Start the server
	log.Info("Dependencies initialized, starting server...", nil)
	server.Serve()
}
