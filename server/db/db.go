package db

import (
	"context"

	"github.com/jmoiron/sqlx"
	_ "github.com/jackc/pgx/stdlib"
)

type database struct {
	db *sqlx.DB
}


func NewDatabase(ctx context.Context, addr string) (*database, error) {
	d, err := sqlx.ConnectContext(ctx, "pgx", addr)
	if err != nil {
		return nil, err
	}

	return &database{
		db: d,
	}, nil
}
