package backtrack

// 回溯树等价于子集问题回溯树
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
		// 枚举从 start 到字符串结尾的所有位置 i
		for i := start; i < len(s); i++ {
			// 对非回文子串剪枝
			if isPalindrome(s[start : i+1]) {
				track = append(track, s[start:i+1])
				// 递归调用 backtrack 处理剩余部分
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
