package web

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/AlekseyZapadovnikov/url-shortner/internal/conf"
	"github.com/go-playground/validator/v10"
)

type ShortenerService interface {
	GiveShortUrl(ctx context.Context, url string) (string, error)
}

type Server struct {
	cfg      *conf.HttpConfig
	router   http.Handler
	log      *slog.Logger
	shortner ShortenerService
	validate *validator.Validate
}

func NewServer(cfg *conf.HttpConfig, router http.Handler, log *slog.Logger, shServ ShortenerService, v *validator.Validate) *Server {
	return &Server{cfg: cfg, router: router, log: log, shortner: shServ, validate: v}
}
