package setup

import (
	"context"
	"github.com/matheus-alvs01dev/go-boilerplate/config"
	"github.com/matheus-alvs01dev/go-boilerplate/internal/adapters/db"
	"github.com/matheus-alvs01dev/go-boilerplate/internal/app/di"
	"github.com/matheus-alvs01dev/go-boilerplate/pkg/log"
	"github.com/pkg/errors"
)

func Setup(ctx context.Context) (*di.Container, error) {
	if err := config.LoadConfig(); err != nil {
		return nil, errors.Wrap(err, "load configs")
	}

	logger, err := log.NewZap(config.GetEnv(), 1)
	if err != nil {
		return nil, errors.Wrap(err, "initialize logger")
	}

	dbClient, err := db.NewClient(config.GetDBConfig().Dsn)
	if err != nil {
		return nil, errors.Wrap(err, "setup database")
	}

	return di.NewContainer(dbClient, logger), nil
}
