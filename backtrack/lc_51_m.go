package backtrack

import "strings"

// 每一层对应一行，每一列对应一层的节点数，遍历该回溯树，对无效的路径进行剪枝
func solveNQueens(n int) [][]string {
	board := make([][]string, n)
	for i := 0; i < n; i++ {
		board[i] = make([]string, n)
		for j := range board[i] {
			board[i][j] = "."
		}
	}
	ans := make([][]string, 0)
	var backtrack func(row int)
	cols := make(map[int]bool)
	up := make(map[int]bool)
	down := make(map[int]bool)
	backtrack = func(row int) {
		if row == n {
			tmp := make([]string, n)
			for i := range board {
				tmp[i] = strings.Join(board[i], "")
			}
			ans = append(ans, tmp)
			return
		}
		for col := range n {
			if cols[col] || up[col+row] || down[col-row] {
				continue
			}
			cols[col] = true
			up[col+row] = true
			down[col-row] = true
			board[row][col] = "Q"
			backtrack(row + 1)
			board[row][col] = "."
			cols[col] = false
			up[col+row] = false
			down[col-row] = false
		}
	}
	backtrack(0)
	return ans
}
