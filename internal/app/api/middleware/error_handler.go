package middleware

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/matheus-alvs01dev/go-boilerplate/internal/app/api/schema"
	"github.com/matheus-alvs01dev/go-boilerplate/pkg/log"
	"github.com/pkg/errors"
	"net/http"
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

func (h *ErrorHandler) Handle(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	statusCode := http.StatusInternalServerError
	httpErr := &HTTPError{
		Message:       "Internal Server Error",
		RequestID:     c.Response().Header().Get(echo.HeaderXRequestID),
		InvalidFields: nil,
	}

	switch {
	case errors.Is(err, sql.ErrNoRows):
		statusCode = http.StatusNotFound
		httpErr.Message = "Resource not found"

	case isEchoHTTPError(err):
		var ee *echo.HTTPError
		errors.As(err, &ee)
		statusCode = ee.Code

		httpErr.Message = ee.Error()

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

	if err := c.JSON(statusCode, httpErr); err != nil {
		h.logger.Error("write out", err, log.Any("err", httpErr))
	}
}

func isEchoHTTPError(err error) bool {
	var ee *echo.HTTPError

	return errors.As(err, &ee)
}

func isSchemaValidationError(err error) bool {
	var ve *schema.ValidationError

	return errors.As(err, &ve)
}
