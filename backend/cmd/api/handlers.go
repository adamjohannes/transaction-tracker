package api

import (
	"encoding/json"
	"errors"
	"monthly-expenses-handler/internal/middleware"
	"net/http"

	"monthly-expenses-handler/internal/api_error"
)

// registerUser
// Handles user registration.
func (app *dependencies) registerUser(w http.ResponseWriter, r *http.Request) {
	var requestBody map[string]any
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	app.logger.Info("Attempting to register a new user", "username", requestBody["username"])

	token, err := app.authController.Register(requestBody)
	if err != nil {
		app.logger.Error("Failed to register user", "error", err.Error())
		var validationErr *apierror.ValidationError
		if errors.As(err, &validationErr) {
			app.respondWithError(w, http.StatusBadRequest, validationErr.Error())
		} else {
			app.respondWithError(w, http.StatusInternalServerError, "Could not register user due to a server error")
		}
		return
	}

	app.logger.Info("Successfully registered user", "username", requestBody["username"])
	app.respondWithJSON(w, http.StatusCreated, map[string]string{"token": token})
}

// loginUser
// Handles user login.
func (app *dependencies) loginUser(w http.ResponseWriter, r *http.Request) {
	var requestBody map[string]any
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	app.logger.Info("Attempting to login user", "username", requestBody["username"])

	token, err := app.authController.Login(requestBody)
	if err != nil {
		app.logger.Error("Failed to login user", "error", err)
		var validationErr *apierror.ValidationError
		if errors.As(err, &validationErr) {
			app.respondWithError(w, http.StatusUnauthorized, validationErr.Error())
		} else {
			app.respondWithError(w, http.StatusInternalServerError, "Could not log in user due to a server error")
		}
		return
	}

	app.logger.Info("Successfully logged in user", "username", requestBody["username"])
	app.respondWithJSON(w, http.StatusOK, map[string]string{"token": token})
}

// createTransaction
// Handles the creation of a new transaction.
func (app *dependencies) createTransaction(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int64)
	if !ok {
		app.respondWithError(w, http.StatusUnauthorized, "Not authorized")
		return
	}

	var requestBody map[string]any
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	app.logger.Info("Attempting to create a new transaction", "userID", userID, "payload", requestBody)

	createdTx, err := app.transactionController.NewTransaction(requestBody, userID)
	if err != nil {
		app.logger.Error("Failed to create transaction", "error", err, "payload", requestBody)
		var validationErr *apierror.ValidationError
		if errors.As(err, &validationErr) {
			app.respondWithError(w, http.StatusBadRequest, validationErr.Error())
		} else {
			app.respondWithError(w, http.StatusInternalServerError, "Could not create transaction due to a server error")
		}
		return
	}

	app.logger.Info("Successfully created transaction", "transaction_id", createdTx.ID, "userID", userID)
	app.respondWithJSON(w, http.StatusCreated, createdTx)
}

// listTransactions
// Handles fetching all transactions.
func (app *dependencies) listTransactions(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int64)
	if !ok {
		app.respondWithError(w, http.StatusUnauthorized, "Not authorized")
		return
	}
	app.logger.Info("Attempting to fetch all transactions", "userID", userID)

	transactions, err := app.transactionController.GetAllTransactionsByUser(userID)
	if err != nil {
		app.logger.Error("Failed to fetch transactions", "error", err, "userID", userID)
		app.respondWithError(w, http.StatusInternalServerError, "Could not retrieve transactions")
		return
	}

	app.logger.Info("Successfully fetched all transactions", "userID", userID, "count", len(transactions))
	app.respondWithJSON(w, http.StatusOK, transactions)
}

// listCategories
// Handles fetching all categories.
func (app *dependencies) listCategories(w http.ResponseWriter, r *http.Request) {
	app.logger.Info("Attempting to fetch all categories")

	categories, err := app.categoryController.GetAllCategories()
	if err != nil {
		app.logger.Error("Failed to fetch categories", "error", err)
		app.respondWithError(w, http.StatusInternalServerError, "Could not retrieve categories")
		return
	}

	app.logger.Info("Successfully fetched all categories", "count", len(categories))
	app.respondWithJSON(w, http.StatusOK, categories)
}

// listAllSubCategories
// Handles fetching all sub-categories.
func (app *dependencies) listAllSubCategories(w http.ResponseWriter, r *http.Request) {
	app.logger.Info("Attempting to fetch all sub categories")

	subCategories, err := app.subCategoryController.GetAllSubCategories()
	if err != nil {
		app.logger.Error("Failed to fetch sub categories", "error", err)
		app.respondWithError(w, http.StatusInternalServerError, "Could not retrieve sub-categories")
		return
	}

	app.logger.Info("Successfully fetched all sub categories", "count", len(subCategories))
	app.respondWithJSON(w, http.StatusOK, subCategories)
}

// listSubCategoriesByCategory
// Handles fetching sub-categories for a parent.
func (app *dependencies) listSubCategoriesByCategory(w http.ResponseWriter, r *http.Request) {
	categoryName := r.PathValue("category_name")
	if categoryName == "" {
		app.respondWithError(w, http.StatusBadRequest, "Category name is required")
		return
	}

	app.logger.Info("Attempting to fetch sub categories related to a parent category", "parent_category", categoryName)

	subCategories, err := app.subCategoryController.GetSubCategoriesByCategory(categoryName)
	if err != nil {
		app.logger.Error("Failed to fetch sub categories", "parent_category", categoryName, "error", err)
		app.respondWithError(w, http.StatusInternalServerError, "Could not retrieve sub-categories")
		return
	}

	app.logger.Info("Successfully fetched sub categories", "parent_category", categoryName, "count", len(subCategories))
	app.respondWithJSON(w, http.StatusOK, subCategories)
}

// listLookups
// is a generic handler for simple lookup tables.
func (app *dependencies) listLookups(fetchFunc func() (any, error), entityName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.logger.Info("Attempting to fetch", "entity", entityName)

		items, err := fetchFunc()
		if err != nil {
			app.logger.Error("Failed to fetch", "entity", entityName, "error", err)
			app.respondWithError(w, http.StatusInternalServerError, "Could not retrieve "+entityName+" list")
			return
		}

		app.logger.Info("Successfully fetched", "entity", entityName)
		app.respondWithJSON(w, http.StatusOK, items)
	}
}
