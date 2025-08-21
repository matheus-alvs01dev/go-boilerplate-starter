package middleware

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/matheus-alvs01dev/go-boilerplate/config"
	"github.com/matheus-alvs01dev/go-boilerplate/internal/app/api/schema"
	"github.com/matheus-alvs01dev/go-boilerplate/pkg/log"
	"github.com/pkg/errors"
)

type HTTPError struct {
	Message       string            `example:"error message"                               json:"message"`
	InvalidFields map[string]string `example:"field: invalid value for this field message" json:"invalid_fields,omitempty"`
	RequestID     string            `example:"nPeca3Cqv9UHYJOZ3NYojBGOFLSVb9zd"            json:"request_id,omitempty"`
}

type ErrorHandler struct {
	logger log.Logger
}

func NewErrorHandler(logger log.Logger) *ErrorHandler {
	return &ErrorHandler{
		logger: logger,
	}
}

// ErrorHandler middleware for Chi router
func (h *ErrorHandler) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				h.handleError(w, r, errors.Errorf("panic: %v", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func (h *ErrorHandler) handleError(w http.ResponseWriter, r *http.Request, err error) {
	statusCode := http.StatusInternalServerError
	httpErr := &HTTPError{
		Message:       err.Error(),
		RequestID:     r.Header.Get("X-Request-ID"),
		InvalidFields: nil,
	}

	switch {
	case errors.Is(err, sql.ErrNoRows):
		statusCode = http.StatusNotFound
		httpErr.Message = "Resource not found"

	case isSchemaValidationError(err):
		var ve *schema.ValidationError
		errors.As(err, &ve)
		statusCode = ve.StatusCode

		httpErr.Message = ve.Error()
		if ve.Field != nil {
			if httpErr.InvalidFields == nil {
				httpErr.InvalidFields = make(map[string]string)
			}

			httpErr.InvalidFields[*ve.Field] = ve.Message
		}
	}

	if config.GetEnv() != "local" && statusCode >= http.StatusInternalServerError {
		h.logger.Error("http error occurred", err)
		
		httpErr.Message = "Internal Server Error"
		httpErr.InvalidFields = nil
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(httpErr); err != nil {
		h.logger.Error("failed to write error response", err, log.Any("err", httpErr))
	}
}

func isSchemaValidationError(err error) bool {
	var ve *schema.ValidationError

	return errors.As(err, &ve)
}
