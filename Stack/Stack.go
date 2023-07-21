package Stack

import "fmt"

type node struct {
	value int
	next  *node
}

type Stack struct {
	head *node
}

func (st *Stack) Push(value int) {
	st.head = &node{
		value: value,
		next:  st.head,
	}
}

func (st *Stack) Pop() (int, error) {
	if st.head == nil {
		return 0, fmt.Errorf("empty stack")
	}
	value := st.head.value
	st.head = st.head.next
	return value, nil
}

func (st *Stack) Top() (int, error) {
	if st.head == nil {
		return 0, fmt.Errorf("empty stack")
	}
	return st.head.value, nil
}
