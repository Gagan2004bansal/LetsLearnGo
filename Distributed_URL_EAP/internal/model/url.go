package model

import "time"

type UrlDB struct {
	Id        string
	Url       string
	Clicked   int64
	ShortCode string
	CreatedAt time.Time
}

type ReqUrl struct {
	Url string `json:"url"`
}

type ResUrl struct {
	Id        string    `json:"id"`
	Url       string    `json:"url"`
	ShortCode string    `json:"shortcode"`
	CreatedAt time.Time `json:"created_at"`
}

func NewShortUrl(Url string, ShortCode string) *UrlDB {
	return &UrlDB{
		Url:       Url,
		ShortCode: ShortCode,
		CreatedAt: time.Now(),
	}
}

func (u *UrlDB) ToResponse() *ResUrl {
	return &ResUrl{
		Id:        u.Id,
		Url:       u.Url,
		ShortCode: u.ShortCode,
		CreatedAt: u.CreatedAt,
	}
}
