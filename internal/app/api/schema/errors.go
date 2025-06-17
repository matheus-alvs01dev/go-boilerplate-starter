package schema

import "net/http"

type ValidationError struct {
	Field      *string `json:"field"`
	Message    string  `json:"message"`
	StatusCode int     `json:"status_code"`
}

func NewValidationError(message string) *ValidationError {
	return &ValidationError{
		Message:    message,
		StatusCode: http.StatusUnprocessableEntity,
	}
}

func (ve ValidationError) WithField(field string) *ValidationError {
	ve.Field = &field

	return &ve
}

func (ve ValidationError) WithStatusCode(statusCode int) *ValidationError {
	ve.StatusCode = statusCode

	return &ve
}

func (ve ValidationError) Error() string {
	if ve.Field != nil {
		return *ve.Field + ": " + ve.Message
	}
	
	return ve.Message
}
