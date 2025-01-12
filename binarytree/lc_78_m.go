package binarytree

func subsets(nums []int) [][]int {
	var backtrack func(s int)
	var track []int
	res := make([][]int, 0)
	backtrack = func(s int) {
		res = append(res, append([]int{}, track...))
		for i := s; i < len(nums); i++ {
			track = append(track, nums[i])
			backtrack(i + 1)
			track = track[:len(track)-1]
		}
	}
	backtrack(0)
	return res
}