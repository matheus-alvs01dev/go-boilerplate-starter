package schema

import (
	"github.com/matheus-alvs01dev/go-boilerplate/pkg/helpers"
	"net/http"
)

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

func (ve ValidationError) WithField(req any, field string) *ValidationError {
	fieldJson := helpers.JSONFieldName(req, field)
	ve.Field = &fieldJson

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
