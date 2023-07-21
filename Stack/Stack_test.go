package Stack

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPop(t *testing.T) {
	tests := []struct {
		name      string
		expected  int
		errString string
	}{
		{
			name:      "pop#1 positive",
			expected:  1,
			errString: "",
		},
		{
			name:      "pop#2 positive",
			expected:  2,
			errString: "",
		},
		{
			name:      "pop#3 negative",
			expected:  0,
			errString: "empty stack",
		},
	}

	st := Stack{}
	st.Push(2)
	st.Push(1)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			value, err := st.Pop()
			if err == nil {
				require.NoError(t, err, "errors are not equal")
				require.Equal(t, test.expected, value, "values are not equal")
				return
			}
			require.EqualError(t, err, test.errString, "errors are not equal")
		})
	}
}

func TestTop(t *testing.T) {
	tests := []struct {
		name      string
		expected  int
		errString string
	}{
		{
			name:      "pop#1 positive",
			expected:  1,
			errString: "",
		},
		{
			name:      "pop#2 positive",
			expected:  2,
			errString: "",
		},
		{
			name:      "pop#3 negative",
			expected:  0,
			errString: "empty stack",
		},
	}

	st := Stack{}
	st.Push(2)
	st.Push(1)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			value, err := st.Top()
			st.Pop()
			if err == nil {
				require.NoError(t, err, "errors are not equal")
				require.Equal(t, test.expected, value, "values are not equal")
				return
			}
			require.EqualError(t, err, test.errString, "errors are not equal")
		})
	}
}
