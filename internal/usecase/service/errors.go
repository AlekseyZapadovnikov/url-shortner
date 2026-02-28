package service

import "errors"

var (
	ErrShortURLNotFound = errors.New("short url not found")
	ErrShortURLExpired  = errors.New("short url expired")
)
