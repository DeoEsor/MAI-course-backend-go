package pet

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func New(ctx context.Context, db *sqlx.DB) (RepositoryI, error) {

	return &repository{
		db: db,
	}, nil
}
