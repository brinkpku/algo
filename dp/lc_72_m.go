package dp

// 自顶向下，递归+备忘录
func minDistance(word1 string, word2 string) int {
	mem := make([][]int, len(word1))
	for i := range mem {
		mem[i] = make([]int, len(word2))
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || j < 0 {
			return max(i, j) + 1
		}
		if mem[i][j] != -1 {
			return mem[i][j]
		}
		if word1[i] == word2[j] {
			mem[i][j] = dfs(i-1, j-1)
		} else {
			mem[i][j] = min(dfs(i-1, j), dfs(i-1, j-1), dfs(i, j-1)) + 1
		}
		return mem[i][j]
	}
	return dfs(len(word1)-1, len(word2)-1)
}

// 自底向上，dp table
func minDistance2(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
		dp[i][0] = i
	}
	for j := range dp[0] {
		dp[0][j] = j
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}

	return dp[len(word1)][len(word2)]
}

// 空间优化
func minDistance3(word1 string, word2 string) int {
	dp := make([]int, len(word2)+1)
	for j := range dp {
		dp[j] = j
	}
	for i := 1; i <= len(word1); i++ {
		leftUp := dp[0]
		dp[0] = i
		for j := 1; j < len(dp); j++ {
			tmp := dp[j]
			if word1[i-1] == word2[j-1] {
				dp[j] = leftUp
			} else {
				dp[j] = min(dp[j-1], tmp, leftUp) + 1
			}
			leftUp = tmp
		}
	}
	return dp[len(word2)]
}
