package repository

import (
	"context"
	"github.com/DeoEsor/MAI-course-backend-go/internal/bootstrap"
	"github.com/DeoEsor/MAI-course-backend-go/internal/config"
	"github.com/DeoEsor/MAI-course-backend-go/internal/repository/pet"
	"github.com/jmoiron/sqlx"
	"github.com/samber/do/v2"
)

func Register(injector do.Injector) error {
	cfg := do.MustInvoke[*config.Config](injector)
	do.Provide(injector, func(injector do.Injector) (*sqlx.DB, error) {
		db, _, err := bootstrap.ConfigureDb(context.Background(), &cfg.DatabaseConfig)
		if err != nil {
			return nil, err
		}
		return db, nil
	})
	do.Provide(injector, pet.New)

	return nil
}
