package dp

// 自顶向下，dfs记忆化搜索
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	mem := make([][]int, len(s1)+1)
	for i := range mem {
		mem[i] = make([]int, len(s2)+1)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}
	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		if i == 0 && j == 0 {
			return true
		}
		if mem[i][j] != -1 {
			return mem[i][j] == 1
		}
		if i == 0 {
			if s2[:j] == s3[:j] {
				mem[i][j] = 1
			} else {
				mem[i][j] = 0
			}
			return mem[i][j] == 1
		}
		if j == 0 {
			if s1[:i] == s3[:i] {
				mem[i][j] = 1
			} else {
				mem[i][j] = 0
			}
			return mem[i][j] == 1
		}
		left := s1[i-1] == s3[i+j-1] && dfs(i-1, j)
		right := s2[j-1] == s3[i+j-1] && dfs(i, j-1)
		if left || right {
			mem[i][j] = 1
		} else {
			mem[i][j] = 0
		}
		return mem[i][j] == 1
	}
	return dfs(len(s1), len(s2))
}

// 递推，自底向上
func isInterleave2(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	dp := make([][]bool, len(s1)+1)
	for i := range dp {
		dp[i] = make([]bool, len(s2)+1)
	}
	dp[0][0] = true
	for i := 1; i <= len(s1); i++ {
		dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]
	}
	for j := 1; j <= len(s2); j++ {
		dp[0][j] = dp[0][j-1] && s2[j-1] == s3[j-1]
	}
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s3[i+j-1] && dp[i-1][j] {
				dp[i][j] = true
			}
			if s2[j-1] == s3[i+j-1] && dp[i][j-1] {
				dp[i][j] = true
			}
		}
	}
	return dp[len(s1)][len(s2)]
}
