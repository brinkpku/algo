package backtrack

// 无重不可重复选排列问题
func permute(nums []int) [][]int {
	var backtrack func()
	track := make([]int, 0)
	res := make([][]int, 0)
	used := make([]bool, len(nums))
	backtrack = func() {
		if len(track) == len(nums) {
			tmp := make([]int, len(track))
			copy(tmp, track)
			res = append(res, tmp)
			return
		}
		for i := range nums {
			if used[i] {
				continue
			}
			track = append(track, nums[i])
			used[i] = true
			backtrack()
			used[i] = false
			track = track[:len(track)-1]
		}
	}
	backtrack()
	return res
}
