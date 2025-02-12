package dp

import "math"

// 自顶向下，递归+备忘录
func numSquares(n int) int {
	downSqrt := int(math.Sqrt(float64(n)))
	mem := make([][]int, downSqrt+1)
	for i := range mem {
		mem[i] = make([]int, n+1)
		for j := range mem[i] {
			if i == 0 {
				mem[i][j] = math.MaxInt
			} else {
				mem[i][j] = -1
			}
		}
	}
	mem[0][0] = 0
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if mem[i][j] != -1 {
			return mem[i][j]
		}
		if j < i*i {
			mem[i][j] = dfs(i-1, j)
		} else {
			mem[i][j] = min(dfs(i-1, j), dfs(i, j-i*i)+1)
		}
		return mem[i][j]
	}
	return dfs(downSqrt, n)
}

// 自底向上，dp table
func numSquares2(n int) int {
	downSqrt := int(math.Sqrt(float64(n)))
	mem := make([][]int, downSqrt+1)
	for i := range mem {
		mem[i] = make([]int, n+1)
		for j := range mem[i] {
			if i == 0 {
				mem[i][j] = math.MaxInt
			} else {
				mem[i][j] = -1
			}
			if j == 0 {
				mem[i][j] = 0
			}
		}
	}
	for i := 1; i <= downSqrt; i++ {
		for j := 1; j <= n; j++ {
			if j < i*i {
				mem[i][j] = mem[i-1][j]
			} else {
				mem[i][j] = min(mem[i-1][j], mem[i][j-i*i]+1)
			}
		}
	}
	return mem[downSqrt][n]
}

// 空间优化
func numSquares3(n int) int {
	downSqrt := int(math.Sqrt(float64(n)))
	mem := make([][]int, downSqrt+1)
	for i := range mem {
		mem[i] = make([]int, n+1)
		for j := range mem[i] {
			if i == 0 {
				mem[i][j] = math.MaxInt
			} else {
				mem[i][j] = -1
			}
			if j == 0 {
				mem[i][j] = 0
			}
		}
	}
	for i := 1; i <= downSqrt; i++ {
		for j := 1; j <= n; j++ {
			if j < i*i {
				mem[i][j] = mem[i-1][j]
			} else {
				mem[i][j] = min(mem[i-1][j], mem[i][j-i*i]+1)
			}
		}
	}
	return mem[downSqrt][n]
}
