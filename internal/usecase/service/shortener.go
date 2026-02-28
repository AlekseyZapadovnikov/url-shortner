package service

import (
	"context"

	"github.com/AlekseyZapadovnikov/url-shortner/internal/conf"
	"github.com/go-playground/validator/v10"
)

type urlRepository interface {
	Save(ctx context.Context, URL string) (int, error)
	GetShortUrl(ctx context.Context, URL string) (string, error)
	GetLongUrl(ctx context.Context, shortURL string) (string, error)
}
type Shortener struct {
	cfg   *conf.ShortenerConfig
	repo  urlRepository
	valid *validator.Validate
}

func NewShortener(cfg *conf.ShortenerConfig, repo urlRepository, vl *validator.Validate) *Shortener {
	return &Shortener{cfg: cfg, repo: repo, valid: vl}
}
