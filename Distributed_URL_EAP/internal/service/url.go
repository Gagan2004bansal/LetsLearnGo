package service

import (
	"crypto/rand"
	"encoding/base64"
	"log/slog"

	"github.com/Gagan2004bansal/LetsLearnGo/internal/model"
)

type UrlService struct {
}

func NewUrlService() *UrlService {
	return &UrlService{}
}

func getshortcode() string {
	randomBytes := make([]byte, 7)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic("getshortcode does not working rn")
	}
	return base64.URLEncoding.EncodeToString(randomBytes)[:7]
}

func (u *UrlService) CreateShortUrl(Url string) *model.ResUrl {
	slog.Info("URL Service - CreateShortUrl")

	shortCode := getshortcode()

	resp := model.NewShortUrl(Url, shortCode)

	return resp.ToResponse()
}
