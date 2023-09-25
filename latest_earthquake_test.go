package gempago

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLatestEarthQuake(t *testing.T) {
	data, err := LatestEarthQuake()

	require.NoError(t, err)
	require.NotEmpty(t, data)
}
