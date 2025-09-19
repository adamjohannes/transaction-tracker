package handlers

import (
	"log"
	"net/http"

	"monthly-expenses-handler/internal/controller/sub_category"
)

type SubCategoryHandler struct {
	controller *sub_category.SubCategoryController
}

func NewSubCategoryHandler(c *sub_category.SubCategoryController) *SubCategoryHandler {
	return &SubCategoryHandler{controller: c}
}

// ListAllSubCategories
// Method: GET /sub-categories
func (h *SubCategoryHandler) ListAllSubCategories(w http.ResponseWriter, r *http.Request) {
	subCategories, err := h.controller.GetAllSubCategories()
	if err != nil {
		log.Printf("Error fetching all sub-categories: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Could not retrieve sub-categories")
		return
	}
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

	subCategories, err := h.controller.GetSubCategoriesByCategory(categoryName)
	if err != nil {
		log.Printf("Error fetching sub-categories for %s: %v", categoryName, err)
		respondWithError(w, http.StatusInternalServerError, "Could not retrieve sub-categories")
		return
	}
	respondWithJSON(w, http.StatusOK, subCategories)
}
