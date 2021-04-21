package app

import (
	"context"

	"github.com/maitesin/yaus/internal/domain"
)

//go:generate moq -out zmock_url_repository_test.go -pkg app_test . URLsRepository

type URLsRepository interface {
	FindByOriginal(ctx context.Context, original string) (domain.URL, error)
	Save(ctx context.Context, url domain.URL) error
}
