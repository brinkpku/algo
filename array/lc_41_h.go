package array

// 原地hash
func firstMissingPositive(nums []int) int {
	for i := range nums {
		idx := nums[i] - 1
		for 0 <= idx && idx < len(nums) && nums[idx] != idx+1 {
			nums[idx], idx = idx+1, nums[idx]-1
		}
	}
	for i := range nums {
		if i != nums[i]-1 {
			return i + 1
		}
	}
	return len(nums) + 1
}
