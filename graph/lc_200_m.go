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

// 广度优先遍历
func numIslands2(grid [][]byte) int {
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}
	ans := 0
	qr, qc := make([]int, 0), make([]int, 0)
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' && !visited[i][j] {
				ans++
				qr = append(qr, i)
				qc = append(qc, j)
				for len(qr) != 0 {
					r, c := qr[0], qc[0]
					visited[r][c] = true
					if 0 < c && grid[r][c-1] == '1' && !visited[r][c-1] {
						qr = append(qr, r)
						qc = append(qc, c-1)
						visited[r][c-1] = true
					}
					if c < len(grid[0])-1 && grid[r][c+1] == '1' && !visited[r][c+1] {
						qr = append(qr, r)
						qc = append(qc, c+1)
						visited[r][c+1] = true
					}
					if 0 < r && grid[r-1][c] == '1' && !visited[r-1][c] {
						qr = append(qr, r-1)
						qc = append(qc, c)
						visited[r-1][c] = true
					}
					if r < len(grid)-1 && grid[r+1][c] == '1' && !visited[r+1][c] {
						qr = append(qr, r+1)
						qc = append(qc, c)
						visited[r+1][c] = true
					}
					qr, qc = qr[1:], qc[1:]
				}
			}
		}
	}
	return ans
}
