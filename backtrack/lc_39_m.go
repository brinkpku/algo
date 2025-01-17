package backtrack

// 回溯，组合/子集问题，无重复可重复选元素
func combinationSum(candidates []int, target int) [][]int {
	var backtrack func(s int)
	sum := 0
	track := make([]int, 0)
	ans := make([][]int, 0)
	backtrack = func(s int) {
		if sum == target {
			tmp := make([]int, len(track))
			copy(tmp, track)
			ans = append(ans, tmp)
			return
		}
		if sum > target {
			return
		}
		for i := s; i < len(candidates); i++ {
			sum += candidates[i]
			track = append(track, candidates[i])
			backtrack(i)
			sum -= candidates[i]
			track = track[:len(track)-1]
		}
	}
	backtrack(0)
	return ans
}
