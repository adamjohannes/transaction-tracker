package handlers

import (
	"log/slog"
	"net/http"

	"monthly-expenses-handler/internal/controller/sub_category"
)

type SubCategoryHandler struct {
	controller *sub_category.Controller
	logger     *slog.Logger
}

func NewSubCategoryHandler(c *sub_category.Controller, l *slog.Logger) *SubCategoryHandler {
	return &SubCategoryHandler{controller: c, logger: l}
}

// ListAllSubCategories
// Method: GET /sub-categories
func (h *SubCategoryHandler) ListAllSubCategories(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Attempting to fetch all sub categories")

	subCategories, err := h.controller.GetAllSubCategories()
	if err != nil {
		h.logger.Error("Failed to fetch sub categories", "error", err)
		respondWithError(w, http.StatusInternalServerError, "Could not retrieve sub-categories")
		return
	}

	h.logger.Info("Successfully fetched all sub categories", "count", len(subCategories))

	respondWithJSON(w, http.StatusOK, subCategories)
}

// ListSubCategoriesByCategory
// Method: GET /categories/{category_name}/sub-categories
func (h *SubCategoryHandler) ListSubCategoriesByCategory(w http.ResponseWriter, r *http.Request) {
	categoryName := r.PathValue("category_name")
	if categoryName == "" {
		respondWithError(w, http.StatusBadRequest, "Category name is required")
		return
	}

	h.logger.Info("Attempting to fetch sub categories related to a parent category", "parent_category", r.PathValue("category_name"))

	subCategories, err := h.controller.GetSubCategoriesByCategory(categoryName)
	if err != nil {
		h.logger.Error("Failed to fetch sub categories", "parent_category", r.PathValue("category_name"), "error", err)
		respondWithError(w, http.StatusInternalServerError, "Could not retrieve sub-categories")
		return
	}

	h.logger.Info("Successfully fetched sub categories", "parent_category", r.PathValue("category_name"), "count", len(subCategories))

	respondWithJSON(w, http.StatusOK, subCategories)
}
