package backtrack

func partition(s string) [][]string {
	var backtrack func(start int)
	ans := make([][]string, 0)
	track := make([]string, 0)
	backtrack = func(start int) {
		if start == len(s) {
			tmp := make([]string, len(track))
			copy(tmp, track)
			ans = append(ans, tmp)
			return
		}
		for i := start; i < len(s); i++ {
			if isPalindrome(s[start : i+1]) {
				track = append(track, s[start:i+1])
				backtrack(i + 1)
				track = track[:len(track)-1]
			}
		}
	}
	backtrack(0)
	return ans
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
