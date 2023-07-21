package leetcode_1614

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func SolutionTest(t testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "test#1",
			input:    "(1+(2*3)+((8)/4))+1",
			expected: 3,
		},
		{
			name:     "test#2",
			input:    "(1)+((2))+(((3)))",
			expected: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expected, Solution(test.input))
		})
	}
}
