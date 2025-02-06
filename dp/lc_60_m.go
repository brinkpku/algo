package dp

import "math"

// 自顶向下，递归+备忘录
func minPathSum(grid [][]int) int {
	mem := make([][]int, len(grid))
	for r := 0; r < len(mem); r++ {
		mem[r] = make([]int, len(grid[0]))
		for c := 0; c < len(mem[0]); c++ {
			mem[r][c] = -1
		}
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i == 0 && j == 0 {
			return grid[i][j]
		}
		if i < 0 || j < 0 {
			return math.MaxInt
		}
		if mem[i][j] != -1 {
			return mem[i][j]
		}
		up := dfs(i-1, j)
		left := dfs(i, j-1)
		mem[i][j] = min(up, left) + grid[i][j]
		return mem[i][j]
	}
	return dfs(len(grid)-1, len(grid[0])-1)
}

// 自底向上，dp table
func minPathSum2(grid [][]int) int {
	dp := make([][]int, len(grid))
	for r := 0; r < len(dp); r++ {
		dp[r] = make([]int, len(grid[0]))
	}
	for r := 0; r < len(dp); r++ {
		for c := 0; c < len(dp[0]); c++ {
			if r == 0 && c == 0 {
				dp[r][c] = grid[r][c]
			} else if r == 0 {
				dp[r][c] = dp[r][c-1] + grid[r][c]
			} else if c == 0 {
				dp[r][c] = dp[r-1][c] + grid[r][c]
			} else {
				dp[r][c] = min(dp[r-1][c], dp[r][c-1]) + grid[r][c]
			}
		}
	}
	return dp[len(grid)-1][len(grid[0])-1]
}
