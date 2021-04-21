package domain_test

import (
	"testing"
	"time"

	"github.com/maitesin/yaus/internal/domain"

	"github.com/stretchr/testify/require"
)

func TestNewURL(t *testing.T) {
	tests := []struct {
		name        string
		original    string
		shortened   string
		expectedURL domain.URL
		expectedErr error
	}{
		{
			name:        "Given a valid URL and a valid shortened, when the NewURL is called, then it creates a valid URL",
			original:    "https://oscarforner.com",
			shortened:   "wololo",
			expectedURL: domain.URL{Original: "https://oscarforner.com", Shortened: "wololo"},
			expectedErr: nil,
		},
		{
			name:        "Given an invalid URL and an invalid shortened, when the NewURL is called, then it returns an OriginalURLInvalidError",
			original:    "",
			shortened:   "",
			expectedURL: domain.URL{},
			expectedErr: domain.NewOriginalURLInvalidError(""),
		},
		{
			name:        "Given a valid URL and an invalid shortened, when the NewURL is called, then it returns an ErrShortenedValueIsEmpty",
			original:    "https://oscarforner.com",
			shortened:   "",
			expectedURL: domain.URL{},
			expectedErr: domain.ErrShortenedValueIsEmpty,
		},
		{
			name:        "Given an invalid URL and a valid shortened, when the NewURL is called, then it returns an OriginalURLInvalidError",
			original:    "wololo",
			shortened:   "123456890",
			expectedURL: domain.URL{},
			expectedErr: domain.NewOriginalURLInvalidError("wololo"),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			start := time.Now().UTC()
			got, err := domain.NewURL(tt.original, tt.shortened)
			end := time.Now().UTC()

			if tt.expectedErr == nil {
				require.NoError(t, err)
				require.Equal(t, tt.expectedURL.Original, got.Original)
				require.Equal(t, tt.expectedURL.Shortened, got.Shortened)
				require.True(t, start.Before(got.CreatedAt) || start.Equal(got.CreatedAt))
				require.True(t, end.After(got.CreatedAt) || end.Equal(got.CreatedAt))
			} else {
				require.ErrorIs(t, err, tt.expectedErr)
			}
		})
	}
}
