package sorting

// func sortArray(nums []int) []int {
// 	mergeSort(nums, 0, len(nums)-1)
// 	return nums
// }

// 二叉树后续遍历
func mergeSort(nums []int, l, r int) {
	if r-l <= 0 {
		return
	}
	m := l + (r-l)/2
	mergeSort(nums, l, m)
	mergeSort(nums, m+1, r)
	merge(nums, l, m, r)
}

func merge(nums []int, l, m, r int) {
	tmp := make([]int, r-l+1)
	p1, p2 := l, m+1
	i := 0
	for p1 <= m && p2 <= r {
		if nums[p1] < nums[p2] {
			tmp[i] = nums[p1]
			p1++
		} else {
			tmp[i] = nums[p2]
			p2++
		}
		i++
	}
	for ; p1 <= m; i, p1 = i+1, p1+1 {
		tmp[i] = nums[p1]
	}
	for ; p2 <= r; i, p2 = i+1, p2+1 {
		tmp[i] = nums[p2]
	}
	for i := range tmp {
		nums[l+i] = tmp[i]
	}
}
