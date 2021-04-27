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

	tests := []struct {
		name                    string
		urlsRepositoryFindURL   domain.URL
		urlsRepositoryFindError error
		query                   app.Query
		expectedResponse        app.QueryResponse
		expectedErr             error
	}{
		{
			name: "",
			query: app.RetrieveURLByShortenedQuery{
				Shortened: validURL.Shortened,
			},
			urlsRepositoryFindURL: validURL,
			expectedResponse:      validURL,
		},
		{
			name:        "",
			query:       &QueryMock{NameFunc: func() string { return "something" }},
			expectedErr: app.InvalidQueryError{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			urlsRepository := &URLsRepositoryMock{
				FindByShortenedFunc: func(context.Context, string) (domain.URL, error) {
					return tt.urlsRepositoryFindURL, tt.urlsRepositoryFindError
				},
			}

			ruh := app.NewRetrieveURLByShortenedHandler(urlsRepository)
			got, err := ruh.Handle(context.Background(), tt.query)
			if err != nil {
				require.ErrorAs(t, err, &tt.expectedErr)
			}
			require.Equal(t, tt.expectedResponse, got)
		})
	}
}

func TestRetrieveURLByOriginalHandler_Handle(t *testing.T) {
	validURL, err := domain.NewURL("https://oscarforner.com", "abcdef")
	require.NoError(t, err)

	tests := []struct {
		name                    string
		urlsRepositoryFindURL   domain.URL
		urlsRepositoryFindError error
		query                   app.Query
		expectedResponse        app.QueryResponse
		expectedErr             error
	}{
		{
			name: "",
			query: app.RetrieveURLByOriginalQuery{
				Original: validURL.Original,
			},
			urlsRepositoryFindURL: validURL,
			expectedResponse:      validURL,
		},
		{
			name:        "",
			query:       &QueryMock{NameFunc: func() string { return "something" }},
			expectedErr: app.InvalidQueryError{},
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
			}

			ruh := app.NewRetrieveURLByOriginalHandler(urlsRepository)
			got, err := ruh.Handle(context.Background(), tt.query)
			if err != nil {
				require.ErrorAs(t, err, &tt.expectedErr)
			}
			require.Equal(t, tt.expectedResponse, got)
		})
	}
}
