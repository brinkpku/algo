package dp

import "math"

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}
	ans := math.MinInt
	for i := range nums {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
	}
	for i := range dp {
		ans = max(dp[i], ans)
	}
	return ans
}
