package middleware

import (
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/matheus-alvs01dev/go-boilerplate/pkg/log"
)

func Logger(logger log.Logger) func(c echo.Context, v echomiddleware.RequestLoggerValues) error {
	return func(c echo.Context, v echomiddleware.RequestLoggerValues) error {
		logger.Info("request",
			log.Any("URI", v.URI),
			log.Any("status", v.Status),
			log.Any("method", v.Method),
			log.Any("remote_ip", v.RemoteIP),
			log.Any("user_agent", v.UserAgent),
			log.Any("latency", v.Latency),
		)

		return nil
	}
}
