package redis

import (
	"context"
	"log/slog"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis() (*redis.Client, error) {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "redis#docker",
		DB:       0,
		Protocol: 2,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		rdb.Close()
		return nil, err
	}

	slog.Info("redis connected...")
	return rdb, nil
}
