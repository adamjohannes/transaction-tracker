package api_error

import "fmt"

// AuthError
// An error type specifically for authentication failures.
// This allows controllers to identify this error and return a 401 Unauthorized.
type AuthError struct {
	Message string
}

func (e *AuthError) Error() string {
	return e.Message
}

func NewAuthError(format string, args ...any) error {
	return &AuthError{
		Message: fmt.Sprintf(format, args...),
	}
}
