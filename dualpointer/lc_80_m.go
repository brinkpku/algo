package dualpointer

func removeDuplicates2(nums []int) int {
	l, r := 1, 1
	count := 0
	maxCount := 2
	for ; r < len(nums); r++ {
		if nums[r] == nums[r-1] {
			count++
			if count < maxCount {
				nums[l] = nums[r]
				l++
			}
		} else {
			count = 0
			nums[l] = nums[r]
			l++
		}
	}
	return l
}
