package liberrors

import (
	"fmt"
)

// Define a custom error type
type KarniError struct {
	Code    int
	Message string
	Err     error
}

func (e *KarniError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("Code: %d, Message: %s, Error: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

// Wrap an existing error with a custom error
func WrapError(code int, message string, err error) *KarniError {
	return &KarniError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}
