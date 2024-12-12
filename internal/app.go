package internal

import (
	"context"
	"errors"
	"github.com/DeoEsor/MAI-course-backend-go/internal/consumer"
	"github.com/DeoEsor/MAI-course-backend-go/internal/controller"
	"github.com/DeoEsor/MAI-course-backend-go/internal/external"
	"github.com/DeoEsor/MAI-course-backend-go/internal/producer"
	"github.com/DeoEsor/MAI-course-backend-go/internal/repository"
	"github.com/DeoEsor/MAI-course-backend-go/internal/service"
	"github.com/DeoEsor/MAI-course-backend-go/internal/use_case"
	log "github.com/rs/zerolog"
	"os"
	"os/signal"
	"syscall"

	"github.com/samber/do/v2"

	"github.com/DeoEsor/MAI-course-backend-go/internal/config"
)

type App struct {
	injector *do.RootScope
}

func New(ctx context.Context, cfg *config.Config) (*App, error) {
	injector := do.New()
	logger := log.New(log.ConsoleWriter{Out: os.Stdout}).
		With().Timestamp().Caller().
		Ctx(ctx).
		Logger()

	do.Provide(injector, cfg.Register)

	if err := errors.Join(
		external.Register(injector),
		producer.Register(injector),
		repository.Register(injector),
		service.Register(injector),
		use_case.Register(injector),
		controller.Register(injector),
		consumer.Register(injector),
	); err != nil {
		logger.Fatal().
			Err(err).
			Msg("failed to register dependencies")
	}

	return &App{
		injector: injector,
	}, nil
}

func (app *App) Run(ctx context.Context) error {
	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

	switch <-signalChannel {
	case os.Interrupt:
		//handle SIGINT
	case syscall.SIGTERM:
		//handle SIGTERM
	}

	return app.injector.ShutdownWithContext(ctx)
}
