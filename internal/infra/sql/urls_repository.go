package sql

import (
	"context"
	"sync"
	"time"

	"github.com/maitesin/yaus/internal/app"
	"github.com/maitesin/yaus/internal/domain"
	"github.com/upper/db/v4"
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

const (
	urlsTable = "urls"
)

type URL struct {
	Original  string    `db:"original"`
	Shortened string    `db:"shortened"`
	CreatedAt time.Time `db:"created_at"`
}

type URLsRepository struct {
	sess db.Session
}

func NewURLsRepository(sess db.Session) *URLsRepository {
	return &URLsRepository{sess: sess}
}

func (ur *URLsRepository) FindByOriginal(ctx context.Context, original string) (domain.URL, error) {
	var result URL

	err := ur.sess.WithContext(ctx).
		Collection(urlsTable).
		Find(db.Cond{"original": original}).
		One(&result)
	if err != nil {
		if err == db.ErrNoMoreRows {
			return domain.URL{}, app.NewURLNotFound(original)
		}
		return domain.URL{}, err
	}

	return domain.URL{
			Original:  result.Original,
			Shortened: result.Shortened,
			CreatedAt: result.CreatedAt,
		},
		nil
}

func (ur *URLsRepository) FindByShortened(ctx context.Context, shortened string) (domain.URL, error) {
	var result URL

	err := ur.sess.WithContext(ctx).
		Collection(urlsTable).
		Find(db.Cond{"shortened": shortened}).
		One(&result)
	if err != nil {
		return domain.URL{}, err
	}

	return domain.URL{
			Original:  result.Original,
			Shortened: result.Shortened,
			CreatedAt: result.CreatedAt,
		},
		nil
}

func (ur *URLsRepository) Save(ctx context.Context, url domain.URL) error {
	dbURL := URL{
		Original:  url.Original,
		Shortened: url.Shortened,
		CreatedAt: url.CreatedAt,
	}

	_, err := ur.sess.WithContext(ctx).
		SQL().
		InsertInto(urlsTable).
		Values(dbURL).
		Exec()

	return err
}
