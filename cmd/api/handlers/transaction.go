package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"monthly-expenses-handler/internal/controller/transaction"
)

// TransactionHandler
// Holds the transaction controller.
type TransactionHandler struct {
	controller *transaction.TransactionController
}

// NewTransactionHandler
// Creates a new handler with the necessary dependencies.
func NewTransactionHandler(c *transaction.TransactionController) *TransactionHandler {
	return &TransactionHandler{controller: c}
}

// CreateTransaction
// Handles the creation of a new transaction.
// Method: POST /transactions
func (h *TransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var requestBody map[string]any
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	createdTx, err := h.controller.NewTransaction(requestBody)
	if err != nil {
		// TODO: Check for validation error (400) or server error (500)
		log.Printf("Error creating transaction: %v", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, createdTx)
}

// ListTransactions
// Handles fetching all transactions.
// Method: GET /transactions
func (h *TransactionHandler) ListTransactions(w http.ResponseWriter, r *http.Request) {
	transactions, err := h.controller.GetAllTransactions()
	if err != nil {
		log.Printf("Error fetching transactions: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Could not retrieve transactions")
		return
	}

	respondWithJSON(w, http.StatusOK, transactions)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload any) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
