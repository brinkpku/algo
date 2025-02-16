package dp

// dp table
func longestPalindrome(s string) string {
	mem := make([][]bool, len(s))
	maxLen := 0
	left, right := 0, 0
	for i := range mem {
		mem[i] = make([]bool, len(s))
	}
	for i := len(mem) - 1; i >= 0; i-- {
		for j := len(mem) - 1; j >= 0 && j >= i; j-- {
			if i == j {
				mem[i][j] = true
			} else if i == j-1 {
				mem[i][j] = s[i] == s[j]
			} else {
				mem[i][j] = mem[i+1][j-1] && s[i] == s[j]
			}
			if mem[i][j] && maxLen < j-i+1 {
				maxLen = j - i + 1
				left, right = i, j
			}
		}
	}
	return s[left : right+1]
}
