package main

import (
	"context"
	_ "github.com/joho/godotenv/autoload"
	"github.com/matheus-alvs01dev/go-boilerplate/config"
	"github.com/matheus-alvs01dev/go-boilerplate/internal/app/api"
	"github.com/matheus-alvs01dev/go-boilerplate/internal/app/setup"
	"github.com/matheus-alvs01dev/go-boilerplate/pkg/log"
	"github.com/pkg/errors"
)

func main() {
	ctx := context.Background()

	if err := run(ctx); err != nil {
		panic(err)
	}
}

func run(ctx context.Context) error {
	if err := config.LoadConfig(); err != nil {
		return errors.Wrap(err, "load configs")
	}

	stp, err := setup.Start(ctx)
	if err != nil {
		return errors.Wrap(err, "start setup")
	}

	stp.Container.Logger().Info(
		"setup components successfully initialized",
		log.Any("env", config.GetEnv()),
	)

	server := api.NewServer(ctx, stp.Container.Logger())
	server.ConfigureRoutes(stp.Container.UserController())

	if err := server.Serve(); err != nil {
		return errors.Wrap(err, "start server")
	}

	return nil
}
