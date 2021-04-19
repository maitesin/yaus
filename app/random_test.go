package app_test

import (
	"testing"
	"time"

	"github.com/maitesin/yaus/app"
	"github.com/stretchr/testify/require"
)

func TestRandomStringGeneratorWithTimeProviderUTC(t *testing.T) {
	t.Parallel()

	timeProvider := &app.TimeProviderUTC{}
	size := 12
	generator := app.NewRandomStringGenerator(timeProvider, size)
	randomString := generator.Generate()

	require.Len(t, randomString, size)
}

func TestRandomStringGeneratorWithUnixEpochTimeProvider(t *testing.T) {
	t.Parallel()

	expectedString := "SRZqe6P56fHw"
	timeProvider := &TimeProviderMock{
		NowFunc: func() time.Time {
			return time.Time{}
		},
	}
	size := 12
	generator := app.NewRandomStringGenerator(timeProvider, size)
	randomString := generator.Generate()

	require.Len(t, randomString, size)
	require.Equal(t, expectedString, randomString)
}
