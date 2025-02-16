package stack

func nextGreaterElements(nums []int) []int {
	stack := []int{}
	ans := make([]int, len(nums))
	for i := range ans {
		ans[i] = -1
	}
	for i := 2*len(nums) - 1; i >= 0; i-- {
		for len(stack) != 0 && nums[i%len(nums)] >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) != 0 {
			ans[i%len(nums)] = stack[len(stack)-1]
		}
		stack = append(stack, nums[i%len(nums)])
	}
	return ans
}
