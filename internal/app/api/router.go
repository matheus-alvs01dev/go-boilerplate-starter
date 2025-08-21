package api

import (
	"net/http"
)

func (s *Server) ConfigureRoutes(
// ..set handlers here
) {
	s.router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "ok"}`))
	})
}
