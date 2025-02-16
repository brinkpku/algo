package dp

// 自底向上，递推方式+空间优化
func uniquePaths(m int, n int) int {
	dp := make([]int, n)
	for i := range m {
		left := dp[0]
		for j := range n {
			if i == 0 || j == 0 {
				dp[j] = 1
			} else {
				dp[j] = left + dp[j]
			}
			left = dp[j]
		}
	}
	return dp[n-1]
}
