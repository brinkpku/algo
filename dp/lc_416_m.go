package dp

// 分割等和子集，自顶向下递归+备忘录
func canPartition(nums []int) bool {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	mem := make([][]int, len(nums))
	for i := range mem {
		mem[i] = make([]int, sum/2)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}
	var dfs func(i, target int) bool
	dfs = func(i, target int) bool {
		if i < 0 || target <= 0 {
			return target == 0
		}
		if mem[i][target-1] != -1 {
			return mem[i][target-1] == 1
		}
		if dfs(i-1, target) || dfs(i-1, target-nums[i]) {
			mem[i][target-1] = 1
		} else {
			mem[i][target-1] = 0
		}
		return mem[i][target-1] == 1
	}
	return dfs(len(nums)-1, sum/2)
}

// 自底向上，dp table
func canPartition2(nums []int) bool {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	dp := make([][]bool, len(nums)+1)
	for i := range dp {
		dp[i] = make([]bool, sum/2+1)
	}
	dp[0][0] = true
	for i := 0; i < len(nums); i++ {
		for j := 0; j <= sum/2; j++ {
			dp[i+1][j] = j >= nums[i] && dp[i][j-nums[i]] || dp[i][j]
		}
	}
	return dp[len(nums)][sum/2]
}
