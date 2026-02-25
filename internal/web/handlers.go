package web

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

// POST /shorten
type PostShortenReq struct {
	RealUrl    string  `json:"real_url" validate:"required,url"`
	ShortenUrl *string `json:"shorten_url"`
}

type postShortenResp struct {
	ShortURL string `json:"short_url"`
}

func (s *Server) handlePostShorten(w http.ResponseWriter, r *http.Request) {
	op := "post.shorten"
	log := s.log
	defer r.Body.Close()

	var req PostShortenReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest,
			getErrResp(ErrorBadRequest, "invalid request"),
		)
		log.ErrorContext(r.Context(), "decode_failed",
			slog.String("op", op),
			slog.Any("err", err),
		)
	}
	if err := s.validate.Struct(req); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			writeJSON(w, http.StatusUnprocessableEntity,
				getErrResp(ErrorInvalidData, "invalid request"),
			)
			return
		}
		log.ErrorContext(r.Context(), "validate_failed", slog.Any("err", err))
		writeJSON(w, http.StatusInternalServerError, getErrResp(ErrorInternal, "internal error"))
		return
	}

	short, err := s.shortner.GiveShortUrl(r.Context(), req.RealUrl)
	if err != nil {
		if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
			writeJSON(w, http.StatusRequestTimeout, getErrResp(ErrorTimeout, "timeout"))
			return
		}
		log.ErrorContext(r.Context(), "shorten_failed",
			slog.String("real_url", req.RealUrl),
			slog.Any("err", err),
		)
		// TODO возможно стоит в дальнейшем добавить подробную обработку ошибок
		writeJSON(w, http.StatusInternalServerError, getErrResp(ErrorInternal, "internal error"))
		return
	}

	writeJSON(w, http.StatusCreated, postShortenResp{ShortURL: short})
}
