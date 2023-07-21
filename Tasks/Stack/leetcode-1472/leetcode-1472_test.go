package leetcode_1472

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSolution(t *testing.T) {
	bh := Constructor("leetcode.com")
	bh.Visit("google.com")
	bh.Visit("facebook.com")
	bh.Visit("youtube.com")
	require.Equal(t, "facebook.com", bh.Back(1))
	require.Equal(t, "google.com", bh.Back(1))
	require.Equal(t, "facebook.com", bh.Forward(1))
	bh.Visit("linkedin.com")
	require.Equal(t, "linkedin.com", bh.Forward(2))
	require.Equal(t, "google.com", bh.Back(2))
	require.Equal(t, "leetcode.com", bh.Back(7))
}
