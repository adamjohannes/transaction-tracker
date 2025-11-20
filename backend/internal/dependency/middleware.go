package dependencies

import (
	"net/http"
)

// listCategories
// Handles fetching all categories.
func (deps *Dependencies) listCategories(w http.ResponseWriter, r *http.Request) {
	deps.Logger.Info("Attempting to fetch all categories...", nil)

	categories, err := deps.CategoryController.GetAllCategories()
	if err != nil {
		deps.Logger.Error("Failed to fetch categories", map[string]interface{}{"error": err})
		deps.respondWithError(w, http.StatusInternalServerError, "Could not retrieve categories")
		return
	}

	deps.Logger.Info("Successfully fetched all categories", map[string]interface{}{"count": len(categories)})
	deps.respondWithJSON(w, http.StatusOK, categories)
}

// listAllSubCategories
// Handles fetching all sub-categories.
func (deps *Dependencies) listAllSubCategories(w http.ResponseWriter, r *http.Request) {
	deps.Logger.Info("Attempting to fetch all sub categories", nil)

	subCategories, err := deps.SubCategoryController.GetAllSubCategories()
	if err != nil {
		deps.Logger.Error("Failed to fetch sub categories", map[string]interface{}{"error": err})
		deps.respondWithError(w, http.StatusInternalServerError, "Could not retrieve sub-categories")
		return
	}

	deps.Logger.Info("Successfully fetched all sub categories", map[string]interface{}{"count": len(subCategories)})
	deps.respondWithJSON(w, http.StatusOK, subCategories)
}

// listSubCategoriesByCategory
// Handles fetching sub-categories for a parent.
func (deps *Dependencies) listSubCategoriesByCategory(w http.ResponseWriter, r *http.Request) {
	categoryName := r.PathValue("category_name")
	if categoryName == "" {
		deps.respondWithError(w, http.StatusBadRequest, "Category name is required")
		return
	}

	deps.Logger.Info("Attempting to fetch sub categories related to a parent category", map[string]interface{}{"parent_category": categoryName})

	subCategories, err := deps.SubCategoryController.GetSubCategoriesByCategory(categoryName)
	if err != nil {
		deps.Logger.Error("Failed to fetch sub categories", map[string]interface{}{"parent_category": categoryName, "error": err})
		deps.respondWithError(w, http.StatusInternalServerError, "Could not retrieve sub-categories")
		return
	}

	deps.Logger.Info("Successfully fetched sub categories", map[string]interface{}{"parent_category": categoryName, "count": len(subCategories)})
	deps.respondWithJSON(w, http.StatusOK, subCategories)
}

// listLookups
// is a generic handler for simple lookup tables.
func (deps *Dependencies) listLookups(fetchFunc func() (any, error), entityName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		deps.Logger.Info("Attempting to fetch", map[string]interface{}{"entity": entityName})

		items, err := fetchFunc()
		if err != nil {
			deps.Logger.Error("Failed to fetch", map[string]interface{}{"entity": entityName, "error": err})
			deps.respondWithError(w, http.StatusInternalServerError, "Could not retrieve "+entityName+" list")
			return
		}

		deps.Logger.Info("Successfully fetched", map[string]interface{}{"entity": entityName})
		deps.respondWithJSON(w, http.StatusOK, items)
	}
}
