package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/matheus-alvs01dev/go-boilerplate/config"
	"github.com/matheus-alvs01dev/go-boilerplate/internal/app/api/middleware"
	"github.com/matheus-alvs01dev/go-boilerplate/pkg/log"
	"github.com/pkg/errors"
)

type Server struct {
	ctx     context.Context
	router  chi.Router
	logger  log.Logger
	apiPort uint16
	server  *http.Server
}

func NewServer(ctx context.Context, logger log.Logger) *Server {
	cfgs := config.GetServerConfig()

	router := chi.NewRouter()

	router.Use(chimiddleware.Recoverer)
	router.Use(middleware.NewErrorHandler(logger).Handle)
	router.Use(middleware.Logger(logger))
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	}))

	svr := &Server{
		ctx:     ctx,
		router:  router,
		logger:  logger,
		apiPort: cfgs.APIPort,
	}

	return svr
}

func (s *Server) Serve() error {
	const timeout = 30 * time.Second

	s.server = &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%d", s.apiPort),
		Handler:      s.router,
		WriteTimeout: timeout,
		ReadTimeout:  timeout,
	}

	serverErr := make(chan error, 1)

	go func() {
		s.logger.Info(fmt.Sprintf("Starting server on port %d...", s.apiPort))
		if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErr <- fmt.Errorf("server failed to start: %w", err)
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-serverErr:
		return err
	case sig := <-signalChan:
		s.logger.Info(fmt.Sprintf("Received signal: %v. Starting graceful shutdown...", sig))
		return s.Shutdown()
	}
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	s.logger.Info("Shutting down server...")

	if err := s.server.Shutdown(ctx); err != nil {
		s.logger.Error("Server shutdown failed", err)

		return fmt.Errorf("server shutdown failed: %w", err)
	}

	s.logger.Info("Server shutdown completed successfully")
	return nil
}
