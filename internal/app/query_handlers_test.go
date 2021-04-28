package app_test

import (
	"context"
	"testing"

	"github.com/maitesin/yaus/internal/app"
	"github.com/maitesin/yaus/internal/domain"
	"github.com/stretchr/testify/require"
)

func TestRetrieveURLByShortenedHandler_Handle(t *testing.T) {
	validURL, err := domain.NewURL("https://oscarforner.com", "abcdef")
	require.NoError(t, err)
	handlerByShortened := func(repository app.URLsRepository) app.QueryHandler {
		return app.NewRetrieveURLByShortenedHandler(repository)
	}
	handlerByOriginal := func(repository app.URLsRepository) app.QueryHandler {
		return app.NewRetrieveURLByOriginalHandler(repository)
	}

	tests := []struct {
		name                    string
		urlsRepositoryFindURL   domain.URL
		urlsRepositoryFindError error
		handlerGenerator        func(repository app.URLsRepository) app.QueryHandler
		query                   app.Query
		expectedResponse        app.QueryResponse
		expectedErr             error
	}{
		{
			name: "",
			query: app.RetrieveURLByShortenedQuery{
				Shortened: validURL.Shortened,
			},
			handlerGenerator:      handlerByShortened,
			urlsRepositoryFindURL: validURL,
			expectedResponse:      validURL,
		},
		{
			name:             "",
			query:            &QueryMock{NameFunc: func() string { return "by shortened" }},
			handlerGenerator: handlerByShortened,
			expectedErr:      app.InvalidQueryError{},
		},
		{
			name: "",
			query: app.RetrieveURLByOriginalQuery{
				Original: validURL.Original,
			},
			handlerGenerator:      handlerByOriginal,
			urlsRepositoryFindURL: validURL,
			expectedResponse:      validURL,
		},
		{
			name:             "",
			query:            &QueryMock{NameFunc: func() string { return "by original" }},
			handlerGenerator: handlerByOriginal,
			expectedErr:      app.InvalidQueryError{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			urlsRepository := &URLsRepositoryMock{
				FindByOriginalFunc: func(context.Context, string) (domain.URL, error) {
					return tt.urlsRepositoryFindURL, tt.urlsRepositoryFindError
				},
				FindByShortenedFunc: func(context.Context, string) (domain.URL, error) {
					return tt.urlsRepositoryFindURL, tt.urlsRepositoryFindError
				},
			}

			ruh := tt.handlerGenerator(urlsRepository)
			got, err := ruh.Handle(context.Background(), tt.query)
			if err != nil {
				require.ErrorAs(t, err, &tt.expectedErr)
			}
			require.Equal(t, tt.expectedResponse, got)
		})
	}
}
