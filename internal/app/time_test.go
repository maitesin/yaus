package app

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTimeProviderUTC(t *testing.T) {
	t.Parallel()

	timeProvider := &TimeProviderUTC{}

	before := time.Now().UTC()
	now := timeProvider.Now()
	after := time.Now().UTC()

	require.True(t, before.Before(now) || before.Equal(now))
	require.True(t, after.After(now) || after.Equal(now))
}
