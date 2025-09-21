package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

// LoggingMiddleware
// Logs the details of each incoming HTTP request.
func LoggingMiddleware(next http.Handler, logger *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Log the incoming request
		logger.Info("Incoming request",
			"method", r.Method,
			"url", r.URL.String(),
			"remote_addr", r.RemoteAddr,
		)

		// Serve the next handler in the chain
		next.ServeHTTP(w, r)

		// Log the completion of the request
		logger.Info("Request completed",
			"method", r.Method,
			"url", r.URL.String(),
			"duration", time.Since(start).String(),
		)
	})
}
