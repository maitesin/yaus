package domain

import (
	"math/rand"
	"time"
)

//go:generate moq -out mock_test.go -pkg domain_test . TimeProvider

// TimeProvider defines the way to obtain the current time
type TimeProvider interface {
	Now() time.Time
}

// TimeProviderUTC provides time in UTC
type TimeProviderUTC struct{}

// Now returns the current time
func (tpu *TimeProviderUTC) Now() time.Time {
	return time.Now().UTC()
}

// URL contains the information for the shortened URL
type URL struct {
	Original  string
	Shortened string
	CreatedAt time.Time
}

// NewURL is a constructor
func NewURL(tp TimeProvider, original string) URL {
	t := tp.Now().UTC()
	return URL{
		Original:  original,
		Shortened: generateRandomString(t.UnixNano(), 12),
		CreatedAt: t,
	}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateRandomString(seed int64, size uint8) string {
	rand.Seed(seed)
	b := make([]byte, size)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
