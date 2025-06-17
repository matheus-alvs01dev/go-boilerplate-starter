package setup

import (
	"context"
	"github.com/matheus-alvs01dev/go-boilerplate/config"
	"github.com/matheus-alvs01dev/go-boilerplate/internal/adapters/db"
	"github.com/matheus-alvs01dev/go-boilerplate/internal/app/di"
	"github.com/matheus-alvs01dev/go-boilerplate/pkg/log"
	"github.com/pkg/errors"
)

type Setup struct {
	Container *di.Container
}

func Start(ctx context.Context) (*Setup, error) {
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

	return &Setup{Container: di.NewContainer(dbClient.DB(), logger)}, nil
}
