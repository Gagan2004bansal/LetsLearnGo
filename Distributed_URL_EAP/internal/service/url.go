package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"log/slog"

	"github.com/Gagan2004bansal/LetsLearnGo/internal/model"
	"github.com/Gagan2004bansal/LetsLearnGo/internal/repository"
)

type UrlService struct {
	repo      repository.UrlRepository
	redisRepo repository.RedisRepository
}

func NewUrlService(repo repository.UrlRepository, redisRepo repository.RedisRepository) *UrlService {
	return &UrlService{
		repo:      repo,
		redisRepo: redisRepo,
	}
}

func getshortcode() string {
	randomBytes := make([]byte, 7)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic("getshortcode does not working rn")
	}
	return base64.URLEncoding.EncodeToString(randomBytes)[:7]
}

func (u *UrlService) CreateShortUrl(ctx context.Context, url string) (*model.ResUrl, error) {
	slog.Info("URL Service - CreateShortUrl")

	for {
		shortCode := getshortcode()

		urlDB := model.NewShortUrl(url, shortCode)

		err := u.repo.Create(ctx, urlDB)

		if err != nil {
			slog.Warn("failed to create URL", "error", err)
			continue
		}

		err = u.redisRepo.SetShortCode(ctx, urlDB)

		if err != nil {
			// If short code collided, we generate another one.
			// For now we retry.
			slog.Warn("failed to create URL, retrying", "error", err)
		}

		return urlDB.ToResponse(), nil
	}
}

func (u *UrlService) GetLongUrl(ctx context.Context, shortCode string) (*model.UrlDB, error) {
	slog.Info("URL Service - GetLongUrl", "shortcode", shortCode)

	// Redis Hit
	url, err := u.redisRepo.GetShortCode(ctx, shortCode)
	if err == nil {
		slog.Info("redis-cache hit", "shortcode", shortCode)
		return url, nil
	}

	// Redis Miss
	if err == repository.ErrURLNotFound {
		slog.Info("redis cache miss", "shortcode", shortCode)
	} else {
		slog.Warn("redis error", "error", err)
	}

	// PostgreSQL
	url, err = u.repo.GetByShortCode(ctx, shortCode)
	if err != nil {
		return nil, err
	}

	// Store Result in Redis
	if err := u.redisRepo.SetShortCode(ctx, url); err != nil {
		slog.Warn("failed to cache url", "error", err)
	}

	return url, nil
}

func (u *UrlService) DeleteShortUrl(
	ctx context.Context,
	shortCode string,
) error {

	slog.Info(
		"URL Service - DeleteShortUrl",
		"shortcode",
		shortCode,
	)

	// Deleted From PostgreSQL
	err := u.repo.DeleteByShortCode(
		ctx,
		shortCode,
	)
	if err != nil {
		return err
	}

	// Delete From Redis
	err = u.redisRepo.DeleteShortCode(ctx, shortCode)
	if err != nil {
		slog.Warn("failed to delete url from cache", "error", err)
	}

	return nil
}

func (u *UrlService) IncrementClicked(
	ctx context.Context,
	shortCode string,
) error {

	return u.repo.IncrementClicked(
		ctx,
		shortCode,
	)
}
