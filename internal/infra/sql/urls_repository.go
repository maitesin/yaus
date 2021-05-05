package sql

import (
	"context"
	"sync"

	"github.com/maitesin/yaus/internal/app"
	"github.com/maitesin/yaus/internal/domain"
)

// InMemoryURLsRepository is a temporary repository until the one that stores in the information in the DB is implemented
type InMemoryURLsRepository struct {
	urls []domain.URL
	m    sync.RWMutex
}

// NewInMemoryURLsRepository is a constructor
func NewInMemoryURLsRepository() *InMemoryURLsRepository {
	// Temporary value added for end-to-end testing purposes
	url, _ := domain.NewURL("https://oscarforner.com", "wololo")
	return &InMemoryURLsRepository{
		urls: []domain.URL{url},
	}
}

// FindByOriginal is self explanatory
func (imur *InMemoryURLsRepository) FindByOriginal(_ context.Context, original string) (domain.URL, error) {
	imur.m.RLock()
	defer imur.m.RUnlock()

	for _, url := range imur.urls {
		if url.Original == original {
			return url, nil
		}
	}
	return domain.URL{}, app.NewURLNotFound(original)
}

// FindByShortened is self explanatory
func (imur *InMemoryURLsRepository) FindByShortened(_ context.Context, shortened string) (domain.URL, error) {
	imur.m.RLock()
	defer imur.m.RUnlock()

	for _, url := range imur.urls {
		if url.Shortened == shortened {
			return url, nil
		}
	}
	return domain.URL{}, app.NewURLNotFound(shortened)
}

// Save is self explanatory
func (imur *InMemoryURLsRepository) Save(_ context.Context, url domain.URL) error {
	imur.m.Lock()
	defer imur.m.Unlock()

	imur.urls = append(imur.urls, url)

	return nil
}
