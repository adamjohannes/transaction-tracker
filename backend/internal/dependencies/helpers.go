package dependencies

import (
	"encoding/json"
	"net/http"
)

// respondWithError
// Sends a JSON error message.
func (deps *Dependencies) respondWithError(w http.ResponseWriter, code int, message string) {
	deps.respondWithJSON(w, code, map[string]string{"error": message})
}

// respondWithJSON
// Writes a JSON response with a given status code and payload.
func (deps *Dependencies) respondWithJSON(w http.ResponseWriter, code int, payload any) {
	response, err := json.Marshal(payload)
	if err != nil {
		deps.Logger.Error("Failed to marshal JSON response", map[string]interface{}{"error": err})
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "Dependencies/json")
	w.WriteHeader(code)
	w.Write(response)
}
