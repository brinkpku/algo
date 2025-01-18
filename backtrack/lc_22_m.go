package backtrack

import "strings"

func generateParenthesis(n int) []string {
	p := []string{"(", ")"}
	var backtrack func()
	track := []string{}
	ans := []string{}
	backtrack = func() {
		if len(track) == 2*n {
			if validParenthesis(track) {
				ans = append(ans, strings.Join(track, ""))
			}
			return
		}
		for i := 0; i < len(p); i++ {
			if len(track) > 0 && string(track[0]) == ")" {
				continue
			}
			track = append(track, string(p[i]))
			backtrack()
			track = track[:len(track)-1]
		}
	}
	backtrack()
	return ans
}

func validParenthesis(s []string) bool {
	leftCount := 0
	for _, r := range s {
		if string(r) == "(" {
			leftCount++
		} else {
			if leftCount == 0 {
				return false
			}
			leftCount--
		}
	}
	return leftCount == 0
}
