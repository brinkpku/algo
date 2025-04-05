package dualpointer

func removeDuplicates(nums []int) int {
	l, r := 0, 0
	for ; r < len(nums); r++ {
		if r > 0 && nums[r] == nums[r-1] {
			continue
		}
		nums[l] = nums[r]
		l++
	}
	return l
}
