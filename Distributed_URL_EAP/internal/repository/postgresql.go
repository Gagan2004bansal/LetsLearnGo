package repository

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect() (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(
		context.Background(),
		"postgres://myuser:distributedSQL@localhost:5433/mydatabase",
	)

	if err != nil {
		return nil, err
	}

	if err := pool.Ping(context.Background()); err != nil {
		pool.Close()
		return nil, err
	}

	slog.Info("postgreSql connection success...")
	return pool, nil
}
