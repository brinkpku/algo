package dp

func trap(height []int) int {
	leftDp := make([]int, len(height))
	rightDp := make([]int, len(height))
	for i := range height {
		if i > 0 {
			leftDp[i] = max(leftDp[i-1], height[i])
		} else {
			leftDp[i] = height[i]
		}
	}
	rightDp[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		rightDp[i] = max(rightDp[i+1], height[i])
	}
	ans := 0
	for i := range height {
		if height[i] < min(leftDp[i], rightDp[i]) {
			ans += min(leftDp[i], rightDp[i]) - height[i]
		}
	}
	return ans
}
