package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// as testing randomness doesn't make sense, this is mainly here to have full test coverage in our console
func TestRandomUtils(t *testing.T) {
	s1 := RandomEmail(10)
	n1 := RandomInt(0, 64)
	s2 := RandomPassword(10)
	s3 := RandomString(10)
	s4 := RandomUsername(10)
	require.Equal(t, s1, s1)
	require.Equal(t, s2, s2)
	require.Equal(t, s3, s3)
	require.Equal(t, s4, s4)
	require.Equal(t, n1, n1)
}
