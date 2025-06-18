package di

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/matheus-alvs01dev/go-boilerplate/internal/adapters/db/repository"
	"github.com/matheus-alvs01dev/go-boilerplate/internal/app/api/ctrl"
	"github.com/matheus-alvs01dev/go-boilerplate/internal/domain/service"
	"github.com/matheus-alvs01dev/go-boilerplate/pkg/log"
)

type Container struct {
	db     *pgxpool.Pool
	logger log.Logger

	userRepo    *repository.UserRepository
	userService *service.UserService
}

func NewContainer(db *pgxpool.Pool, logger log.Logger) *Container {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	return &Container{
		logger:      logger,
		db:          db,
		userRepo:    userRepo,
		userService: userService,
	}
}

func (c *Container) DB() *pgxpool.Pool {
	return c.db
}

func (c *Container) Logger() log.Logger {
	return c.logger
}

func (c *Container) UserController() *ctrl.UserController {
	return ctrl.NewUserController(c.userService)
}
