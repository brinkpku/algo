package dp

import "math"

// 自顶向下递归+备忘录
func coinChange(coins []int, amount int) int {
	var dp func(amount int) int
	mem := make([]int, amount+1)
	for i := range mem {
		mem[i] = -999
	}
	dp = func(amount int) int {
		if amount == 0 {
			return 0
		}
		if amount < 0 {
			return -1
		}
		if mem[amount] != -999 {
			return mem[amount]
		}
		ans := math.MaxInt
		for i := range coins {
			num := dp(amount - coins[i])
			if num == -1 {
				continue
			}
			ans = min(num+1, ans)
		}
		if ans == math.MaxInt {
			ans = -1
		}
		mem[amount] = ans
		return ans
	}
	return dp(amount)
}

// 自底向上，dp table
func coinChange2(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, c := range coins {
			if i-c < 0 {
				continue
			}
			dp[i] = min(dp[i-c]+1, dp[i])
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
