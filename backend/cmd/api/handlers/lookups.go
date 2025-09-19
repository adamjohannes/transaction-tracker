package handlers

import (
	"log"
	"monthly-expenses-handler/internal/controller/currency"
	"monthly-expenses-handler/internal/controller/status"
	"monthly-expenses-handler/internal/controller/transaction_type"
	"net/http"
)

type StatusHandler struct {
	controller *status.StatusController
}

func NewStatusHandler(c *status.StatusController) *StatusHandler {
	return &StatusHandler{controller: c}
}

func (h *StatusHandler) ListStatus(w http.ResponseWriter, r *http.Request) {
	items, err := h.controller.GetAllStatus()
	if err != nil {
		log.Printf("Error fetching status list: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Could not retrieve status list")
		return
	}
	respondWithJSON(w, http.StatusOK, items)
}

type CurrencyHandler struct {
	controller *currency.CurrencyController
}

func NewCurrencyHandler(c *currency.CurrencyController) *CurrencyHandler {
	return &CurrencyHandler{controller: c}
}

func (h *CurrencyHandler) ListCurrencies(w http.ResponseWriter, r *http.Request) {
	items, err := h.controller.GetAllCurrencies()
	if err != nil {
		log.Printf("Error fetching currency list: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Could not retrieve currency list")
		return
	}
	respondWithJSON(w, http.StatusOK, items)
}

type TypeHandler struct {
	controller *transaction_type.TypeController
}

func NewTypeHandler(c *transaction_type.TypeController) *TypeHandler {
	return &TypeHandler{controller: c}
}

func (h *TypeHandler) ListTypes(w http.ResponseWriter, r *http.Request) {
	items, err := h.controller.GetAllTransactionTypes()
	if err != nil {
		log.Printf("Error fetching type list: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Could not retrieve type list")
		return
	}
	respondWithJSON(w, http.StatusOK, items)
}
