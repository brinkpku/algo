package dp

// 自顶向下，递归+备忘录
func rob(nums []int) int {
	var dfs func(i int) int
	mem := make([]int, len(nums))
	for i := range mem {
		mem[i] = -1
	}
	dfs = func(i int) int {
		if mem[i] != -1 {
			return mem[i]
		}
		if i == 0 {
			mem[i] = nums[i]
		} else if i == 1 {
			mem[i] = max(nums[0], nums[1])
		} else {
			mem[i] = max(dfs(i-2)+nums[i], dfs(i-1))
		}
		return mem[i]
	}
	return dfs(len(nums) - 1)
}
