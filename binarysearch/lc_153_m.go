package binarysearch

// 确定循环不变量，处理边界条件
func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] > nums[len(nums)-1] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return nums[l]
}
