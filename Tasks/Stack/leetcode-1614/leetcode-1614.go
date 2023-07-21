package leetcode_1614

func Solution(s string) int {
	max, count := 0, 0

	for _, c := range s {
		if c == '(' {
			count++
			if count > max {
				max = count
			}
			continue
		}
		if c == ')' {
			count--
			continue
		}
	}

	return max
}
