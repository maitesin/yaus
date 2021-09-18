package sql

import (
	"context"
	"time"

	"github.com/maitesin/yaus/internal/app"
	"github.com/maitesin/yaus/internal/domain"
	"github.com/upper/db/v4"
)

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
