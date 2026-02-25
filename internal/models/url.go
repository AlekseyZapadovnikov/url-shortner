package models

type Url string

type ShortenResult struct {
	ShortURL    Url `validate:"required,url"`
	OriginalURL Url `validate:"required,url"`
}
