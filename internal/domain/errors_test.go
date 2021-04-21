package domain_test

import (
	"testing"

	"github.com/maitesin/yaus/internal/domain"

	"github.com/stretchr/testify/require"
)

func TestNewOriginalURLInvalidError(t *testing.T) {
	t.Parallel()

	url := "wololo"
	err := domain.NewOriginalURLInvalidError(url)

	require.Equal(t, `invalid URL "wololo"`, err.Error())
}
