package api

import (
	"context"
	"log"
	"net/http"

	"monthly-expenses-handler/cmd/api/handlers"
	"monthly-expenses-handler/internal/controller/transaction"

	"github.com/jackc/pgx/v5/pgxpool"
)

// StartServer
// Initializes and runs the headless API server.
func StartServer(pool *pgxpool.Pool) {
	ctx := context.Background()
	transactionController := transaction.NewTransactionController(pool, ctx)

	transactionHandler := handlers.NewTransactionHandler(transactionController)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /transactions", transactionHandler.CreateTransaction)
	mux.HandleFunc("GET /transactions", transactionHandler.ListTransactions)

	port := "8080"
	log.Printf("🚀 Starting API server on http://localhost:%s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("❌ Could not start server: %s\n", err)
	}
}
