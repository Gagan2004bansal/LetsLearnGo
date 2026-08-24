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
	repo repository.UrlRepository
}

func NewUrlService(repo repository.UrlRepository) *UrlService {
	return &UrlService{
		repo: repo,
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

		if err == nil {
			return urlDB.ToResponse(), nil
		}

		// If short code collided, we generate another one.
		// For now we retry.
		slog.Warn(
			"failed to create URL, retrying",
			"error", err,
		)
	}
}

func (u *UrlService) GetLongUrl(ctx context.Context, shortCode string) (*model.UrlDB, error) {
	slog.Info(
		"URL Service - GetLongUrl",
		"shortcode",
		shortCode,
	)

	url, err := u.repo.GetByShortCode(
		ctx,
		shortCode,
	)

	if err != nil {
		return nil, err
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

	return u.repo.DeleteByShortCode(
		ctx,
		shortCode,
	)
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
