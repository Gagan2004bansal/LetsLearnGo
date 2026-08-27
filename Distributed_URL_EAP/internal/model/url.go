package model

import (
	"time"

	"github.com/google/uuid"
)

// Tell about schema that our database have
type UrlDB struct {
	Id        string
	Url       string
	Clicked   int64
	ShortCode string
	CreatedAt time.Time
}

// Tell about schema that our request comes like these
type ReqUrl struct {
	Url string `json:"url"`
}

// Tell about schema that what we send in response
type ResUrl struct {
	Id        string    `json:"id"`
	Url       string    `json:"url"`
	ShortCode string    `json:"shortcode"`
	CreatedAt time.Time `json:"created_at"`
}

// Function to create a record with defined database schema
func NewShortUrl(url string, shortCode string) *UrlDB {
	return &UrlDB{
		Id:        uuid.New().String(),
		Url:       url,
		Clicked:   0,
		ShortCode: shortCode,
		CreatedAt: time.Now(),
	}
}

// Function to send response with reponse schema
func (u *UrlDB) ToResponse() *ResUrl {
	return &ResUrl{
		Id:        u.Id,
		Url:       u.Url,
		ShortCode: u.ShortCode,
		CreatedAt: u.CreatedAt,
	}
}
