package logger

import (
	"log/slog"
	"strings"
	"time"

	"github.com/AlekseyZapadovnikov/url-shortner/internal/conf"
)

const (
	KeyRequestID = "request_id"
	KeyUserAgent = "user_agent"
	KeyMethod    = "method"
	KeyPath      = "path"
	KeyStatus    = "status"
	KeyDuration  = "duration_ms"
	KeyShortURL  = "short_url"
	KeyOriginal  = "original_url"
)

func New(cfg *conf.LoggerConfig) *slog.Logger {
	opts := &slog.HandlerOptions{
		Level:     cfg.Level,
		AddSource: cfg.AddSource,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				t := a.Value.Time()
				if cfg.TimeFormat != "" {
					return slog.String(a.Key, t.Format(cfg.TimeFormat))
				}
				return slog.String(a.Key, t.Format(time.RFC3339Nano))
			}

			switch strings.ToLower(a.Key) {
			case "authorization", "password", "passwd", "token",
				"access_token", "refresh_token", "api_key", "secret":
				return slog.String(a.Key, "***")
			case "cookie", "set-cookie":
				return slog.Attr{}
			}

			return a
		},
	}
	var h slog.Handler
	switch cfg.Format {
	case conf.FormatText:
		h = slog.NewTextHandler(cfg.Output, opts)
	case conf.FormatJSON:
		h = slog.NewJSONHandler(cfg.Output, opts)
	default:
		h = slog.NewTextHandler(cfg.Output, opts)
	}
	// Базовые поля сервиса — всегда в каждом логе
	base := slog.New(h).With(
		slog.String("app", cfg.AppName),
		slog.String("env", cfg.Env),
	)

	return base
}

func SetDefault(l *slog.Logger) {
	slog.SetDefault(l)
}
