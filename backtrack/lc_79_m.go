package backtrack

func exist(board [][]byte, word string) bool {
	var backtrack func(row, col, k int)
	ans := false
	backtrack = func(row, col, k int) {
		if k == len(word) {
			ans = true
			return
		}
		if row < 0 || row >= len(board) ||
			col >= len(board[0]) || col < 0 {
			return
		}
		if board[row][col] != word[k] || board[row][col] == 0 {
			return
		}
		board[row][col] = 0
		backtrack(row-1, col, k+1)
		backtrack(row+1, col, k+1)
		backtrack(row, col-1, k+1)
		backtrack(row, col+1, k+1)
		board[row][col] = word[k]
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			backtrack(i, j, 0)
			if ans {
				return ans
			}
		}
	}
	return ans
}
