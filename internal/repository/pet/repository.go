package pet

import (
	"context"
	"github.com/samber/do/v2"
	log "github.com/sirupsen/logrus"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func New(injector do.Injector) (RepositoryI, error) {
	log.WithFields(log.Fields{
		"service": "repository",
	}).
		Info("service invoked")
	return &repository{
		db: do.MustInvoke[*sqlx.DB](injector),
	}, nil
}

func (repository *repository) Shutdown(ctx context.Context) error {
	log.WithContext(ctx).
		WithFields(log.Fields{
			"service": "repository",
		}).
		Info("service shutdown")
	return repository.db.Close()
}
