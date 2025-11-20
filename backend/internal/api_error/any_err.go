package api_error

import "fmt"

// AnyError
// A generic error type.
// This allows controllers to identify this error and return a 500 Internal Server Error.
type AnyError struct {
	Message string
}

func (e *AnyError) Error() string {
	return e.Message
}

func NewAnyError(format string, args ...any) error {
	return &AnyError{
		Message: fmt.Sprintf(format, args...),
	}
}
