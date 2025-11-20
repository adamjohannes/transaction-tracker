package handlers

import (
	"log/slog"
	"net/http"

	"monthly-expenses-handler/internal/controller/category"
)

type CategoryHandler struct {
	controller *category.Controller
	logger     *slog.Logger
}

func NewCategoryHandler(c *category.Controller, l *slog.Logger) *CategoryHandler {
	return &CategoryHandler{controller: c, logger: l}
}

func (h *CategoryHandler) ListCategories(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Attempting to fetch all categories")

	categories, err := h.controller.GetAllCategories()
	if err != nil {
		h.logger.Error("Failed to fetch categories", "error", err)
		respondWithError(w, http.StatusInternalServerError, "Could not retrieve categories")
		return
	}

	h.logger.Info("Successfully fetched all categories", "count", len(categories))

	respondWithJSON(w, http.StatusOK, categories)
}
