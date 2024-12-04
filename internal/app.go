package internal

import (
	"context"
	"errors"
	"github.com/DeoEsor/MAI-course-backend-go/internal/repository"
	"github.com/DeoEsor/MAI-course-backend-go/internal/repository/pet"
	log "github.com/sirupsen/logrus"

	"github.com/samber/do/v2"

	"github.com/DeoEsor/MAI-course-backend-go/internal/config"
)

type App struct {
	injector *do.RootScope
}

func New(ctx context.Context, cfg *config.Config) (*App, error) {
	injector := do.New()
	do.Provide(injector, cfg.Register)
	err := errors.Join(
		repository.Register(injector),
		// use_case.Register ...
	)
	if err != nil {
		log.Fatalf("failed to register dependencies: %v", err)
	}
	do.MustInvoke[pet.RepositoryI](injector)
	return &App{
		injector: injector,
	}, nil
}

func (app *App) Run(ctx context.Context) error {
	err := app.injector.ShutdownWithContext(ctx)
	if err != nil {
		log.Fatalf("failed to shutdown application: %v", err)
	}
	return nil
}
