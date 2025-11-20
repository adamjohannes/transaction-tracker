package dependencies

import (
	"net/http"
)

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
