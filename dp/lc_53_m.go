package dp

import "math"

// 自底向上，dp table，最优结果不一定在最后
func maxSubArray(nums []int) int {
	ans := math.MinInt
	dp := make([]int, len(nums))
	for i := range nums {
		if i == 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = max(dp[i-1]+nums[i], nums[i])
		}
		ans = max(dp[i], ans)
	}
	return ans
}

// 空间优化
func maxSubArray2(nums []int) int {
	ans := math.MinInt
	prev := 0
	for i := range nums {
		prev = max(prev+nums[i], nums[i])
		ans = max(prev, ans)
	}
	return ans
}
