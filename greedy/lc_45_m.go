package greedy

import "math"

// 动态规划做法，自顶向下，递归+备忘录
func jump(nums []int) int {
	var dfs func(i int) int
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = -1
	}
	dfs = func(i int) int {
		if i == 0 {
			return 0
		}
		if dp[i] != -1 {
			return dp[i]
		}
		ans := math.MaxInt
		for j := i - 1; j >= 0; j-- {
			if nums[j]+j >= i {
				ans = min(ans, dfs(j)+1)
			}
		}
		dp[i] = ans
		return ans
	}
	return dfs(len(nums) - 1)
}

// 动态规划做法，自底向上
func jump2(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = math.MaxInt
	}
	dp[0] = 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if j+nums[j] >= i {
				dp[i] = min(dp[j]+1, dp[i])
			}
		}
	}
	return dp[len(nums)-1]
}

// 贪心做法
func jump3(nums []int) int {
	ans := 0
	maxJump := 0
	end := 0
	for i := range nums {
		if end >= len(nums)-1 {
			break
		}
		maxJump = max(nums[i]+i, maxJump)
		if i == end {
			ans++
			end = maxJump
		}
	}
	return ans
}
