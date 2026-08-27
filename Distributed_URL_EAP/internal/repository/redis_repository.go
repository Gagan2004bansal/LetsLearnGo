package repository

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/Gagan2004bansal/LetsLearnGo/internal/model"
	"github.com/redis/go-redis/v9"
)

// Interface of cache for abstraction
type RedisRepository interface {
	GetShortCode(ctx context.Context, shortCose string) (*model.UrlDB, error)
	SetShortCode(ctx context.Context, url *model.UrlDB) error
	DeleteShortCode(ctx context.Context, shortCode string) error
}

type RedisUrlRepository struct {
	rdb *redis.Client
}

// Constructor to initialize the redis instance
func NewRedisUrlRepository(rdb *redis.Client) *RedisUrlRepository {
	return &RedisUrlRepository{
		rdb: rdb,
	}
}

// Function to get shortcode in cache
func (r *RedisUrlRepository) GetShortCode(ctx context.Context, shortCode string) (*model.UrlDB, error) {
	value, err := r.rdb.Get(
		ctx,
		shortCode,
	).Result()

	if errors.Is(err, redis.Nil) {
		return nil, ErrURLNotFound
	}

	if err != nil {
		return nil, err
	}

	var url model.UrlDB

	err = json.Unmarshal(
		[]byte(value),
		&url,
	)

	if err != nil {
		return nil, err
	}

	return &url, nil
}

// Function to set shortcode in cache
func (r *RedisUrlRepository) SetShortCode(ctx context.Context, url *model.UrlDB) error {
	value, err := json.Marshal(url)
	if err != nil {
		return err
	}

	return r.rdb.Set(ctx, url.ShortCode, value, 0).Err()
}

// Function to delete shortcode in cache
func (r *RedisUrlRepository) DeleteShortCode(ctx context.Context, shortCode string) error {
	return r.rdb.Del(ctx, shortCode).Err()
}
