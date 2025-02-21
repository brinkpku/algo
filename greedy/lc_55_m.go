package greedy

func canJump(nums []int) bool {
	farest := 0
	for i := range nums {
		if i <= farest {
			farest = max(nums[i]+i, farest)
			if farest >= len(nums)-1 {
				return true
			}
		}
	}
	return false
}
