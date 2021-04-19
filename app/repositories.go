package app

import (
	"context"

	"github.com/maitesin/yaus/internal/domain"
)

type URLsRepository interface {
	FindByOriginal(ctx context.Context, original string) (domain.URL, error)
	Save(ctx context.Context, url domain.URL) error
}
