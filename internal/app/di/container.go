package di

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/matheus-alvs01dev/go-boilerplate/pkg/log"
)

type Container struct {
	db     *pgxpool.Pool
	logger log.Logger

	// added services and repositories can be added here
}

func NewContainer(db *pgxpool.Pool, logger log.Logger) *Container {
	return &Container{
		logger: logger,
		db:     db,
	}
}
