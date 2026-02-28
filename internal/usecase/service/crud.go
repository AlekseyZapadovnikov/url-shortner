package service

import "context"

func (sh *Shortener) GiveShortUrl(ctx context.Context, realUrl, wantedDomain string) (string, error) {
	// нужна ли валидация?
	shortUrl, err := sh.giveShortUrl(wantedDomain)
}
