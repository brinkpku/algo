package dp

// 自顶向下，递归+备忘录
func wordBreak(s string, wordDict []string) bool {
	hash := make(map[string]bool)
	for _, w := range wordDict {
		hash[w] = true
	}
	mem := make([]int, len(s))
	for i := range mem {
		mem[i] = -1
	}
	var dfs func(end int) bool
	dfs = func(end int) bool {
		if end == 0 {
			return true
		}
		if mem[end-1] != -1 {
			return mem[end-1] == 1
		}
		for i := end - 1; i >= 0; i-- {
			if hash[s[i:end]] && dfs(i) {
				mem[end-1] = 1
				return true
			}
		}
		mem[end-1] = 0
		return false
	}
	return dfs(len(s))
}

// 自底向上，dp table
func wordBreak2(s string, wordDict []string) bool {
	hash := make(map[string]bool)
	for _, w := range wordDict {
		hash[w] = true
	}
	dp := make([]bool, len(s)+1)
	for i := range dp {
		if i == 0 {
			dp[i] = true
		} else {
			for j := i; j >= 0; j-- {
				if dp[j] && hash[s[j:i]] {
					dp[i] = true
					break
				}
			}
		}
	}
	return dp[len(s)]
}
