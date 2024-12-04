package bootstrap

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"

	"github.com/DeoEsor/MAI-course-backend-go/internal/config"
)

func ConfigureDb(ctx context.Context, dbConfig *config.DatabaseConfig) (*sqlx.DB, context.CancelFunc, error) {
	connConfig, err := pgx.ParseConfig(dbConfig.DSN)
	if err != nil {
		return nil, nil, fmt.Errorf("bootstrap.db.init.pgx: failed to parse config, err: %w", err)
	}

	db := sqlx.NewDb(
		stdlib.OpenDB(*connConfig),
		"pgx",
	)

	db.SetMaxOpenConns(dbConfig.DatabaseMaxConnections)

	_, cancel := context.WithTimeout(ctx, dbConfig.DatabaseConnectionTimeout)
	return db, cancel, nil
}
