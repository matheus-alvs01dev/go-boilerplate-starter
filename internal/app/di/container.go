package di

import (
	"database/sql"
	"github.com/matheus-alvs01dev/go-boilerplate/pkg/log"
)

type Container struct {
	db     *sql.DB
	logger log.Logger

	// added services and repositories can be added here
}

func NewContainer(db *sql.DB, logger log.Logger) *Container {
	return &Container{
		logger: logger,
		db:     db,
	}
}
