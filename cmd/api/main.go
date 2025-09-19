package main

import (
	"context"
	"log"
	"net/http"

	"monthly-expenses-handler/cmd/api/handlers"
	"monthly-expenses-handler/internal/controller/transaction"
	"monthly-expenses-handler/internal/database"
)

func main() {
	pool := database.ConnectDB()
	defer pool.Close()

	ctx := context.Background()
	transactionController := transaction.NewTransactionController(pool, ctx)

	transactionHandler := handlers.NewTransactionHandler(transactionController)

	// Set up the router
	mux := http.NewServeMux()
	mux.HandleFunc("POST /transactions", transactionHandler.CreateTransaction)
	mux.HandleFunc("GET /transactions", transactionHandler.ListTransactions)

	// Start the HTTP server
	port := "8080"
	log.Printf("🚀 Starting API server on http://localhost:%s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("❌ Could not start server: %s\n", err)
	}
}
