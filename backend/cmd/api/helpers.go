package api

import (
	"encoding/json"
	"net/http"
)

// respondWithError
// Sends a JSON error message.
func (app *dependencies) respondWithError(w http.ResponseWriter, code int, message string) {
	app.respondWithJSON(w, code, map[string]string{"error": message})
}

// respondWithJSON
// Writes a JSON response with a given status code and payload.
func (app *dependencies) respondWithJSON(w http.ResponseWriter, code int, payload any) {
	response, err := json.Marshal(payload)
	if err != nil {
		app.logger.Error("Failed to marshal JSON response", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "dependencies/json")
	w.WriteHeader(code)
	w.Write(response)
}
