package stack

import "fmt"

type node struct {
	value any
	next  *node
}

type Stack struct {
	head *node
}

func New() *Stack {
	return &Stack{}
}

func (st *Stack) Push(v any) {
	st.head = &node{
		value: v,
		next:  st.head,
	}
}

func (st *Stack) Pop() (any, error) {
	if st.head != nil {
		v := st.head.value
		st.head = st.head.next
		return v, nil
	}
	return nil, fmt.Errorf("empty stack")
}

func (st *Stack) Top() (any, error) {
	if st.head != nil {
		return st.head.value, nil
	}
	return nil, fmt.Errorf("empty stack")
}
