package service

import (
	"crypto/rand"
	"encoding/base64"
	"log/slog"
	"sync"

	"github.com/Gagan2004bansal/LetsLearnGo/internal/model"
)

type UrlService struct {
	urls map[string]model.UrlDB
	mu   sync.RWMutex
}

func NewUrlService() *UrlService {
	return &UrlService{
		urls: make(map[string]model.UrlDB),
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

func (u *UrlService) CreateShortUrl(Url string) *model.ResUrl {
	slog.Info("URL Service - CreateShortUrl")

	shortCode := getshortcode()

	resp := model.NewShortUrl(Url, shortCode)

	u.mu.Lock()
	u.urls[shortCode] = *resp
	u.mu.Unlock()

	return resp.ToResponse()
}

func (u *UrlService) GetLongUrl(ShortCode string) (model.UrlDB, bool) {
	u.mu.RLock()
	defer u.mu.RUnlock()

	url, exists := u.urls[ShortCode]
	return url, exists
}
