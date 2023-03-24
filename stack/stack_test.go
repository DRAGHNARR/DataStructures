package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	st := New()

	st.Push(2)
	st.Push(3)

	tests := []struct {
		name string
		want any
		err  error
	}{
		{
			name: "test #1: top/pop exists",
			want: 3,
		},
		{
			name: "test #2: top/pop exists",
			want: 2,
		},
		{
			name: "test #3: top/pop not exists",
			want: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if v, err := st.Top(); v != test.want {
				t.Errorf("Top() = %v, want %v, err: %v", v, test.want, err)
			}
			if v, err := st.Pop(); v != test.want {
				t.Errorf("Pop() = %v, want %v, err: %v", v, test.want, err)
			}
		})
	}
}
