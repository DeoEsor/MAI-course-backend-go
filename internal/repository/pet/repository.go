package pet

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func New(ctx context.Context, db *sqlx.DB) (RepositoryI, error) {

	return &Repository{
		db: db,
	}, nil
}
