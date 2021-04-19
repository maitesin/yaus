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
	generator := app.NewRandomStringGenerator(timeProvider)
	size := 12
	randomString := generator.Generate(size)

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
	generator := app.NewRandomStringGenerator(timeProvider)
	size := 12
	randomString := generator.Generate(size)

	require.Equal(t, expectedString, randomString)
}
