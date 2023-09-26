package gempago

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAffectedEarthQuake(t *testing.T) {
	data, err := AffectedEarthQuake()

	require.NoError(t, err)
	require.NotEmpty(t, data)
}
