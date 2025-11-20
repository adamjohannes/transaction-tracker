package middleware

import (
	"context"
	"encoding/json"
	"monthly-expenses-handler/internal/service/auth"
	"net/http"
	"strings"
)

type contextKey string

const UserIDKey contextKey = "userID"
const GroupBy contextKey = "groupBy"

// AuthMiddleware
// Validates the JWT token from the Authorization header.
func AuthMiddleware(next http.Handler, authSvc *auth.AuthService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			respondWithError(w, http.StatusUnauthorized, "Authorization header required")
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			respondWithError(w, http.StatusUnauthorized, "Invalid Authorization header format")
			return
		}

		tokenString := parts[1]
		userID, err := authSvc.ValidateJWT(tokenString)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "Invalid or expired token")
			return
		}

		// Add userID to the request context
		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// respondWithError
// A helper to send JSON error messages from middleware.
func respondWithError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
