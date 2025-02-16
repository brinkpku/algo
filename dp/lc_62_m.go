package dp

// 自底向上，递推方式+空间优化
func uniquePaths(m int, n int) int {
	dp := make([]int, n)
	for i := range m {
		for j := range n {
			if i == 0 || j == 0 {
				dp[j] = 1
			} else {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[n-1]
}
