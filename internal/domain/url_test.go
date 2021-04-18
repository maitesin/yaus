package domain_test

import (
	"testing"
	"time"

	"github.com/maitesin/yaus/internal/domain"
	"github.com/stretchr/testify/require"
)

func TestNewURL(t *testing.T) {
	tests := []struct {
		name         string
		timeProvider domain.TimeProvider
		url          string
		expected     domain.URL
	}{
		{
			name: "",
			timeProvider: &TimeProviderMock{
				NowFunc: func() time.Time {
					return time.Time{}
				},
			},
			url: "https://oscarforner.com",
			expected: domain.URL{
				Shortened: "SRZqe6P56fHw",
				Original:  "https://oscarforner.com",
				CreatedAt: time.Time{},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := domain.NewURL(tt.timeProvider, tt.url)
			require.Equal(t, tt.expected, got)
		})
	}
}
