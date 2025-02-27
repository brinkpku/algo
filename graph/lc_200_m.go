package graph

// 深度优先遍历
func numIslands(grid [][]byte) int {
	var dfs func(i, j int)
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}
	dfs = func(i, j int) {
		if i < 0 || j < 0 || i == len(grid) || j == len(grid[0]) {
			return
		}
		if visited[i][j] {
			return
		}
		visited[i][j] = true
		if grid[i][j] == '1' {
			dfs(i+1, j)
			dfs(i-1, j)
			dfs(i, j+1)
			dfs(i, j-1)
		}
	}
	ans := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' && !visited[i][j] {
				dfs(i, j)
				ans++
			}
		}
	}
	return ans
}
