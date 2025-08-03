package middleware

import (
	"net/http"
	"time"

	"github.com/matheus-alvs01dev/go-boilerplate/pkg/log"
)

func Logger(logger log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			
			// Create a custom response writer to capture the status code
			wrappedWriter := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
			
			next.ServeHTTP(wrappedWriter, r)
			
			duration := time.Since(start)
			
			logger.Info("request",
				log.Any("URI", r.RequestURI),
				log.Any("status", wrappedWriter.statusCode),
				log.Any("method", r.Method),
				log.Any("remote_ip", r.RemoteAddr),
				log.Any("user_agent", r.UserAgent()),
				log.Any("latency", duration),
			)
		})
	}
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
