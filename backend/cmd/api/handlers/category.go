package handlers

import (
	"log"
	"net/http"

	"monthly-expenses-handler/internal/controller/category"
)

type CategoryHandler struct {
	controller *category.CategoryController
}

func NewCategoryHandler(c *category.CategoryController) *CategoryHandler {
	return &CategoryHandler{controller: c}
}

func (h *CategoryHandler) ListCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.controller.GetAllCategories()
	if err != nil {
		log.Printf("Error fetching categories: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Could not retrieve categories")
		return
	}
	respondWithJSON(w, http.StatusOK, categories)
}
