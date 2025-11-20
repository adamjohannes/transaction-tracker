package dependencies

import (
	"encoding/json"
	"errors"
	"monthly-expenses-handler/internal/middleware"
	"net/http"

	"monthly-expenses-handler/internal/api_error"
)

// registerUser
// Handles user registration.
func (deps *Dependencies) registerUser(w http.ResponseWriter, r *http.Request) {
	var requestBody map[string]any
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		deps.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	deps.Logger.Info("Attempting to register a new user", map[string]interface{}{"username": requestBody["username"]})

	token, err := deps.AuthController.Register(requestBody)
	if err != nil {
		deps.Logger.Error("Failed to register user", map[string]interface{}{"error": err.Error()})
		var validationErr *apierror.ValidationError
		if errors.As(err, &validationErr) {
			deps.respondWithError(w, http.StatusBadRequest, validationErr.Error())
		} else {
			deps.respondWithError(w, http.StatusInternalServerError, "Could not register user due to a server error")
		}
		return
	}

	deps.Logger.Info("Successfully registered user", map[string]interface{}{"username": requestBody["username"]})
	deps.respondWithJSON(w, http.StatusCreated, map[string]string{"token": token})
}

// loginUser
// Handles user login.
func (deps *Dependencies) loginUser(w http.ResponseWriter, r *http.Request) {
	var requestBody map[string]any
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		deps.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	deps.Logger.Info("Attempting to login user", map[string]interface{}{"username": requestBody["username"]})

	token, err := deps.AuthController.Login(requestBody)
	if err != nil {
		deps.Logger.Error("Failed to login user", map[string]interface{}{"error": err})
		var validationErr *apierror.ValidationError
		if errors.As(err, &validationErr) {
			deps.respondWithError(w, http.StatusUnauthorized, validationErr.Error())
		} else {
			deps.respondWithError(w, http.StatusInternalServerError, "Could not log in user due to a server error")
		}
		return
	}

	deps.Logger.Info("Successfully logged in user", map[string]interface{}{"username": requestBody["username"]})
	deps.respondWithJSON(w, http.StatusOK, map[string]string{"token": token})
}

// createTransaction
// Handles the creation of a new transaction.
func (deps *Dependencies) createTransaction(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int64)
	if !ok {
		deps.respondWithError(w, http.StatusUnauthorized, "Not authorized")
		return
	}

	var requestBody map[string]any
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		deps.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	deps.Logger.Info("Attempting to create a new transaction", map[string]interface{}{"userID": userID, "payload": requestBody})

	createdTx, err := deps.TransactionController.NewTransaction(requestBody, userID)
	if err != nil {
		deps.Logger.Error("Failed to create transaction", map[string]interface{}{"error": err, "payload": requestBody})
		var validationErr *apierror.ValidationError
		if errors.As(err, &validationErr) {
			deps.respondWithError(w, http.StatusBadRequest, validationErr.Error())
		} else {
			deps.respondWithError(w, http.StatusInternalServerError, "Could not create transaction due to a server error")
		}
		return
	}

	deps.Logger.Info("Successfully created transaction", map[string]interface{}{"transaction_id": createdTx.ID, "userID": userID})
	deps.respondWithJSON(w, http.StatusCreated, createdTx)
}

// listTransactions
// Handles fetching all transactions.
func (deps *Dependencies) listTransactions(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int64)
	if !ok {
		deps.respondWithError(w, http.StatusUnauthorized, "Not authorized")
		return
	}
	deps.Logger.Info("Attempting to fetch all transactions", map[string]interface{}{"userID": userID})

	transactions, err := deps.TransactionController.GetAllTransactionsByUser(userID)
	if err != nil {
		deps.Logger.Error("Failed to fetch transactions", map[string]interface{}{"error": err, "userID": userID})
		deps.respondWithError(w, http.StatusInternalServerError, "Could not retrieve transactions")
		return
	}

	deps.Logger.Info("Successfully fetched all transactions", map[string]interface{}{"userID": userID, "count": len(transactions)})
	deps.respondWithJSON(w, http.StatusOK, transactions)
}

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
