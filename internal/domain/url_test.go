package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewURL(t *testing.T) {
	tests := []struct {
		name        string
		original    string
		shortened   string
		expectedURL URL
		expectedErr error
	}{
		{
			name:        "Given a valid URL and a valid shortened, when the NewURL is called, then it creates a valid URL",
			original:    "https://oscarforner.com",
			shortened:   "wololo",
			expectedURL: URL{Original: "https://oscarforner.com", Shortened: "wololo"},
			expectedErr: nil,
		},
		{
			name:        "Given an invalid URL and an invalid shortened, when the NewURL is called, then it returns an OriginalURLInvalidError",
			original:    "",
			shortened:   "",
			expectedURL: URL{},
			expectedErr: NewOriginalURLInvalidError(""),
		},
		{
			name:        "Given a valid URL and an invalid shortened, when the NewURL is called, then it returns an ShortenedValueIsEmptyError",
			original:    "https://oscarforner.com",
			shortened:   "",
			expectedURL: URL{},
			expectedErr: ShortenedValueIsEmptyError{},
		},
		{
			name:        "Given an invalid URL and a valid shortened, when the NewURL is called, then it returns an OriginalURLInvalidError",
			original:    "wololo",
			shortened:   "123456890",
			expectedURL: URL{},
			expectedErr: NewOriginalURLInvalidError("wololo"),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			start := time.Now().UTC()
			got, err := NewURL(tt.original, tt.shortened)
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
