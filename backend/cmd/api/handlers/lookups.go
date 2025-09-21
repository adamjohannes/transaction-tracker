package handlers

import (
	"log/slog"
	"monthly-expenses-handler/internal/controller/currency"
	"monthly-expenses-handler/internal/controller/status"
	"monthly-expenses-handler/internal/controller/transaction_type"
	"net/http"
)

type StatusHandler struct {
	controller *status.StatusController
	logger     *slog.Logger
}

func NewStatusHandler(c *status.StatusController, l *slog.Logger) *StatusHandler {
	return &StatusHandler{controller: c, logger: l}
}

func (h *StatusHandler) ListStatus(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Attempting to fetch status")

	items, err := h.controller.GetAllStatus()
	if err != nil {
		h.logger.Error("Failed to fetch status", "error", err)
		respondWithError(w, http.StatusInternalServerError, "Could not retrieve status list")
		return
	}

	h.logger.Info("Successfully fetched status", "count", len(items))

	respondWithJSON(w, http.StatusOK, items)
}

type CurrencyHandler struct {
	controller *currency.CurrencyController
	logger     *slog.Logger
}

func NewCurrencyHandler(c *currency.CurrencyController, l *slog.Logger) *CurrencyHandler {
	return &CurrencyHandler{controller: c, logger: l}
}

func (h *CurrencyHandler) ListCurrencies(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Attempting to fetch currencies")

	items, err := h.controller.GetAllCurrencies()
	if err != nil {
		h.logger.Error("Failed to fetch currencies", "error", err)
		respondWithError(w, http.StatusInternalServerError, "Could not retrieve currency list")
		return
	}

	h.logger.Info("Successfully fetched currencies", "count", len(items))

	respondWithJSON(w, http.StatusOK, items)
}

type TypeHandler struct {
	controller *transaction_type.TypeController
	logger     *slog.Logger
}

func NewTypeHandler(c *transaction_type.TypeController, l *slog.Logger) *TypeHandler {
	return &TypeHandler{controller: c, logger: l}
}

func (h *TypeHandler) ListTypes(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Attempting to fetch types")

	items, err := h.controller.GetAllTransactionTypes()
	if err != nil {
		h.logger.Error("Failed to fetch types", "error", err)
		respondWithError(w, http.StatusInternalServerError, "Could not retrieve type list")
		return
	}

	h.logger.Info("Successfully fetched types", "count", len(items))

	respondWithJSON(w, http.StatusOK, items)
}
