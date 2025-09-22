package apierror

import "fmt"

// ValidationError
// An error type specifically for user input validation failures.
// This allows handlers to identify this error and return a 400 Bad Request.
type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

func NewValidationError(format string, args ...any) error {
	return &ValidationError{
		Message: fmt.Sprintf(format, args...),
	}
}
