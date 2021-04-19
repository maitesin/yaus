package domain

import (
	"time"
)

// URL contains the information for the shortened URL
type URL struct {
	Original  string
	Shortened string
	CreatedAt time.Time
}

// NewURL is a constructor
func NewURL(original, shortened string) URL {
	return URL{
		Original:  original,
		Shortened: shortened,
		CreatedAt: time.Now().UTC(),
	}
}
