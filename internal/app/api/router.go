package api

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/matheus-alvs01dev/go-boilerplate/config"
	"github.com/matheus-alvs01dev/go-boilerplate/internal/app/api/ctrl"
	"github.com/matheus-alvs01dev/go-boilerplate/internal/app/api/middleware"
	"github.com/matheus-alvs01dev/go-boilerplate/pkg/log"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

type Server struct {
	ctx     context.Context
	router  *echo.Echo
	logger  log.Logger
	apiPort uint16
}

func NewServer(ctx context.Context, logger log.Logger) *Server {
	cfgs := config.GetServerConfig()
	svr := &Server{
		ctx:     ctx,
		router:  echo.New(),
		logger:  logger,
		apiPort: cfgs.APIPort,
	}

	svr.router.HideBanner = true
	svr.router.Use(echomiddleware.Recover())
	svr.router.HTTPErrorHandler = middleware.NewErrorHandler(logger).Handle
	svr.router.Use(echomiddleware.RequestLoggerWithConfig(echomiddleware.RequestLoggerConfig{
		LogURI:        true,
		LogStatus:     true,
		LogValuesFunc: middleware.Logger(svr.logger),
	}))

	svr.router.Use(echomiddleware.CORSWithConfig(echomiddleware.CORSConfig{
		AllowOrigins: []string{"*"}, // Allow all origins for development purposes
	}))

	return svr
}

func (s *Server) Serve() error {
	const timeout = 30 * time.Second

	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%d", s.apiPort),
		Handler:      s.router,
		WriteTimeout: timeout,
		ReadTimeout:  timeout,
	}

	go func() {
		s.logger.Info("Starting server...")
		if err := s.router.Start(srv.Addr); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return
		}
	}()

	<-s.ctx.Done()

	if err := s.Shutdown(); err != nil {
		return err
	}

	s.logger.Info("Server stopped gracefully")

	return nil
}

func (s *Server) Shutdown() error {
	s.logger.Info("Shutdown command received, shutting down server...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := s.router.Shutdown(shutdownCtx); err != nil {
		return err
	}

	s.logger.Info("Server shutdown gracefully")

	return nil
}

func (s *Server) ConfigureRoutes(
	userController *ctrl.UserController,
) {
	s.router.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	userGroup := s.router.Group("/users")
	userGroup.POST("", userController.Create)
	userGroup.GET("/:id", userController.GetByID)
	userGroup.PUT("/:id", userController.Update)
	userGroup.DELETE("/:id", userController.Delete)
}
