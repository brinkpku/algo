package dualpointer

// 283. 移动零

// 双指针，s指向非零元素应该放置的位置，f指向当前遍历到的位置
func moveZeroes(nums []int) {
	s, f := 0, 0
	for ; f < len(nums); f++ {
		if nums[f] != 0 {
			nums[s] = nums[f]
			s++
		}
	}
	for ; s < len(nums); s++ {
		nums[s] = 0
	}
}
