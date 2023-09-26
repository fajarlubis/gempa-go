package gempago

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRecentsEarthQuake(t *testing.T) {
	data, err := RecentsEarthQuake()

	require.NoError(t, err)
	require.NotEmpty(t, data)
}
