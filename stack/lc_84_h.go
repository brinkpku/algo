package stack

// 单调递增栈找左右第一个更小的边界，使用两种遍历方式。
// 1. 与更小的数反向遍历
// 2. 与更小的数同向遍历
func largestRectangleArea(heights []int) int {
	leftStack := []int{}
	left := make([]int, len(heights))
	for i := range heights {
		for len(leftStack) != 0 && heights[leftStack[len(leftStack)-1]] > heights[i] {
			leftStack = leftStack[:len(leftStack)-1]
		}
		if len(leftStack) != 0 {
			left[i] = leftStack[len(leftStack)-1]
		} else {
			left[i] = -1
		}
		leftStack = append(leftStack, i)
	}
	rightStack := []int{}
	right := make([]int, len(heights))
	for i := range heights {
		for len(rightStack) != 0 && heights[rightStack[len(rightStack)-1]] > heights[i] {
			right[rightStack[len(rightStack)-1]] = i
			rightStack = rightStack[:len(rightStack)-1]
		}
		rightStack = append(rightStack, i)
	}
	for _, i := range rightStack {
		right[i] = len(heights)
	}

	ans := 0
	for i := range heights {
		width := right[i] - left[i] - 1
		ans = max(ans, heights[i]*width)
	}
	return ans
}
