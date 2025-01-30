package binarysearch

func searchMatrix(matrix [][]int, target int) bool {
	h, w := len(matrix), len(matrix[0])
	left, right := 0, h*w-1
	for left <= right {
		mid := left + (right-left)/2
		v := getval(mid, matrix)
		if v == target {
			return true
		} else if v > target {
			right = mid - 1
		} else if v < target {
			left = mid + 1
		}
	}
	return false
}

func getval(id int, m [][]int) int {
	r := id / len(m[0])
	c := id % len(m[0])
	return m[r][c]
}
