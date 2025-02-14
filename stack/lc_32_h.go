package stack

func longestValidParentheses(s string) int {
	stack := []int{-1}
	ans := 0
	for idx, r := range s {
		if r == '(' {
			stack = append(stack, idx)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) != 0 {
				ans = max(ans, idx-stack[len(stack)-1])
			} else {
				stack = append(stack, idx)
			}
		}
	}
	return ans
}
