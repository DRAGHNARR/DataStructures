//package leetcode_84
package main

import (
	"../../../Stack"
	"fmt"
)

func largestRectangleArea(heights []int) int {
	st := &Stack.StackSlice{}
	var largest int

	for index, height := range heights {
		start := index

		for stored, err := st.Top(); err == nil && stored[1] > height; stored, err = st.Top() {
			st.Pop()
			current := stored[1] * (index - stored[0])
			if current > largest {
				largest = current
			}
			start = stored[0]
		}

		st.Push([]int{start, height})
	}

	for stored, err := st.Pop(); err == nil; stored, err = st.Pop() {
		current := stored[1] * (len(heights) - stored[0])
		if current > largest {
			largest = current
		}
	}

	return largest
}

func main() {
	fmt.Println(largestRectangleArea([]int{6, 2, 5, 4, 5, 1, 6}))
}
