package domain

import (
	"net/url"
	"time"
)

// URL contains the information for the shortened URL
type URL struct {
	Original  string
	Shortened string
	CreatedAt time.Time
}

// NewURL is a constructor
func NewURL(original, shortened string) (URL, error) {
	if original == "" {
		return URL{}, NewOriginalURLInvalidError(original)
	}
	if shortened == "" {
		return URL{}, ShortenedValueIsEmptyError{}
	}
	_, err := url.ParseRequestURI(original)
	if err != nil {
		return URL{}, NewOriginalURLInvalidError(original)
	}
	return URL{
		Original:  original,
		Shortened: shortened,
		CreatedAt: time.Now().UTC(),
	}, nil
}
