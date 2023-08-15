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

func (st *Stack) IsEmpty() bool {
	if st.head == nil {
		return true
	}
	return false
}

type nodeSlice struct {
	value []int
	next  *nodeSlice
}

type StackSlice struct {
	head *nodeSlice
}

func (st *StackSlice) Push(value []int) {
	st.head = &nodeSlice{
		value: value,
		next:  st.head,
	}
}

func (st *StackSlice) Pop() ([]int, error) {
	if st.head == nil {
		return nil, fmt.Errorf("empty stack")
	}
	value := st.head.value
	st.head = st.head.next
	return value, nil
}

func (st *StackSlice) Top() ([]int, error) {
	if st.head == nil {
		return nil, fmt.Errorf("empty stack")
	}
	return st.head.value, nil
}

func (st *StackSlice) IsEmpty() bool {
	if st.head == nil {
		return true
	}
	return false
}
