package dp

import "math"

// 自顶向下，递归+备忘录
func maxProfit(prices []int) int {
	var dfs func(i, hold int) int
	mem := make([][2]int, len(prices))
	for i := range mem {
		mem[i] = [2]int{math.MinInt, math.MinInt}
	}
	dfs = func(i, hold int) int {
		if i == 0 {
			if hold == 1 {
				return -prices[i]
			}
			return 0
		}
		if mem[i][hold] != math.MinInt {
			return mem[i][hold]
		}
		if hold == 1 {
			mem[i][1] = max(dfs(i-1, 1), dfs(i-1, 0)-prices[i])
		} else {
			mem[i][0] = max(dfs(i-1, 0), dfs(i-1, 1)+prices[i])
		}
		return mem[i][hold]
	}
	return dfs(len(prices)-1, 0)
}

// 自底向上，递推 dp table
func maxProfit2(prices []int) int {
	mem := make([][2]int, len(prices))
	mem[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		mem[i][0] = max(mem[i-1][1]+prices[i], mem[i-1][0])
		mem[i][1] = max(mem[i-1][1], mem[i-1][0]-prices[i])
	}
	return mem[len(prices)-1][0]
}

// 空间优化
func maxProfit3(prices []int) int {
	f0, f1 := 0, -prices[0]
	tmp := f0
	for i := 1; i < len(prices); i++ {
		f0 = max(f1+prices[i], f0)
		f1 = max(f1, tmp-prices[i])
		tmp = f0
	}
	return f0
}
