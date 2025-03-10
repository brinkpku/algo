package array

func searchMatrix(matrix [][]int, target int) bool {
	r := 0
	c := len(matrix[0]) - 1
	for r < len(matrix) && c >= 0 {
		if matrix[r][c] == target {
			return true
		} else if matrix[r][c] > target {
			c--
		} else if matrix[r][c] < target {
			r++
		}
	}
	return false
}
