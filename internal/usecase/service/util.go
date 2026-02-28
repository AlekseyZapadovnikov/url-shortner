package service

import (
	"context"
	"errors"
	"fmt"
	"math/rand"

	"github.com/AlekseyZapadovnikov/url-shortner/internal/repo"
)

func giveShortUrl(ctx context.Context, wantedDomain string) (string, error) {
	if wantedDomain == "" {
		wantedDomain, err := generateUniqueWantedDomain(ctx)
	}
}

func (sh *Shortener) generateUniqueWantedDomain(ctx context.Context) (string, error) {
	alphabet := sh.cfg.GetAlphabet()
	length := sh.cfg.GetAlphabet()

	res := ""
	for range length {
		res += string(alphabet[rand.Intn(len(alphabet))])
	}

	for {
		_, err := sh.repo.GetLongUrl(context.TODO(), res)
		if err != nil {
			if errors.Is(err, repo.ErrNoRows) {
				return res, nil
			} else {
				return "", fmt.Errorf("GetLongUrl failed: %w", err)
			}
		}
		log.WARN("Коллизия") // !!! я ещё не решил как я буду передавать логгер, как зависимость или через context
	}
}
