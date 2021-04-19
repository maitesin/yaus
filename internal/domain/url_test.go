package domain_test

import (
	"testing"
	"time"

	"github.com/maitesin/yaus/internal/domain"
	"github.com/stretchr/testify/require"
)

func TestNewURL(t *testing.T) {
	t.Parallel()

	url := "https://oscarforner.com"
	shortened := "SRZqe6P56fHw"
	expectedURL := domain.URL{
		Shortened: "SRZqe6P56fHw",
		Original:  "https://oscarforner.com",
	}

	before := time.Now().UTC()
	got := domain.NewURL(url, shortened)
	after := time.Now().UTC()

	require.Equal(t, expectedURL.Original, got.Original)
	require.Equal(t, expectedURL.Shortened, got.Shortened)
	require.True(t, before.Before(got.CreatedAt) || before.Equal(got.CreatedAt))
	require.True(t, after.After(got.CreatedAt) || after.Equal(got.CreatedAt))
}
