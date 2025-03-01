package graph

// bfs遍历
func orangesRotting(grid [][]int) int {
	q := [][]int{}
	fresh := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 2 {
				q = append(q, []int{i, j})
			} else if grid[i][j] == 1 {
				fresh++
			}
		}
	}
	if fresh == 0 {
		return 0
	}
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	ans := -1
	for len(q) > 0 {
		ans++
		for sz := len(q); sz > 0; sz-- {
			node := q[0]
			q = q[1:]
			for _, dir := range directions {
				x := node[0] + dir[0]
				y := node[1] + dir[1]
				if 0 <= x && x < len(grid) && 0 <= y && y < len(grid[0]) &&
					grid[x][y] == 1 {
					q = append(q, []int{x, y})
					grid[x][y] = 2
				}
			}
		}
	}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}
	return ans
}
