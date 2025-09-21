package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"monthly-expenses-handler/internal/controller/transaction"
)

// TransactionHandler
// Holds the transaction controller.
type TransactionHandler struct {
	controller *transaction.TransactionController
	logger     *slog.Logger
}

// NewTransactionHandler
// Creates a new handler with the necessary dependencies.
func NewTransactionHandler(c *transaction.TransactionController, l *slog.Logger) *TransactionHandler {
	return &TransactionHandler{controller: c, logger: l}
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

	h.logger.Info("Attempting to create a new transaction", "payload", requestBody)

	createdTx, err := h.controller.NewTransaction(requestBody)
	if err != nil {
		// TODO: Check for validation error (400) or server error (500)
		h.logger.Error("Failed to create transaction", "error", err, "payload", requestBody)
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	h.logger.Info("Successfully created transaction", "transaction_id", createdTx.ID)

	respondWithJSON(w, http.StatusCreated, createdTx)
}

// ListTransactions
// Handles fetching all transactions.
// Method: GET /transactions
func (h *TransactionHandler) ListTransactions(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Attempting to fetch all transactions")

	transactions, err := h.controller.GetAllTransactions()
	if err != nil {
		h.logger.Error("Failed to fetch transactions", "error", err)
		respondWithError(w, http.StatusInternalServerError, "Could not retrieve transactions")
		return
	}

	h.logger.Info("Successfully fetched all transactions", "count", len(transactions))

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
