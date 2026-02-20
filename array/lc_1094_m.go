package array

/*
差分数组，用来计算区间增量，最后通过前缀和计算出每个位置的最终值
*/
func carPooling(trips [][]int, capacity int) bool {
	diff := make([]int, 1001)
	for _, trip := range trips {
		diff[trip[1]] += trip[0]
		diff[trip[2]] -= trip[0]
	}
	if diff[0] > capacity {
		return false
	}
	for i := 1; i < len(diff); i++ {
		diff[i] += diff[i-1]
		if diff[i] > capacity {
			return false
		}
	}
	return true
}
