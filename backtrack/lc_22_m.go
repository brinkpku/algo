package backtrack

import "strings"

// 无重可重复选组合问题
func generateParenthesis(n int) []string {
	p := []string{"(", ")"}
	var backtrack func()
	track := []string{}
	ans := []string{}
	left, right := 0, 0
	backtrack = func() {
		if left < right || len(track) == 2*n {
			if left == right {
				ans = append(ans, strings.Join(track, ""))
			}
			return
		}
		for i := 0; i < len(p); i++ {
			if p[i] == "(" {
				left++
			} else {
				right++
			}
			track = append(track, string(p[i]))
			backtrack()
			if p[i] == "(" {
				left--
			} else {
				right--
			}
			track = track[:len(track)-1]
		}
	}
	backtrack()
	return ans
}
