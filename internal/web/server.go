package web

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/AlekseyZapadovnikov/url-shortner/internal/conf"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

// мне не нравится, что я url передаю string
type ShortenerService interface {
	GiveShortUrl(ctx context.Context, realUrl, wantedDomain string) (string, error)
	GetRealUrl(ctx context.Context, url string) (string, error)
}

type Server struct {
	cfg *conf.HttpConfig

	shortner ShortenerService

	router *chi.Mux
	server *http.Server

	validate *validator.Validate
	log      *slog.Logger
}

func NewServer(
	cfg *conf.HttpConfig,
	log *slog.Logger,
	shServ ShortenerService,
	v *validator.Validate,
) *Server {
	router := chi.NewRouter()
	adr := cfg.GetAdress()
	server := &http.Server{
		Addr:    adr,
		Handler: router,
	}
	srv := &Server{
		cfg:      cfg,
		router:   router,
		log:      log,
		shortner: shServ,
		validate: v,
		server:   server,
	}

	return srv
}
