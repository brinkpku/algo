package array

import (
	"math/rand"
)

func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, k)
}

func quickSelect(nums []int, l, r, k int) int {
	lt, eq := partition(nums, l, r)
	if lt <= len(nums)-k && len(nums)-k < eq {
		return nums[lt]
	} else if eq <= len(nums)-k {
		return quickSelect(nums, eq, r, k)
	} else {
		return quickSelect(nums, l, lt-1, k)
	}
}

func partition(nums []int, l, r int) (int, int) {
	randomIndex := rand.Intn(r-l+1) + l
	pivot := nums[randomIndex]
	lt, eq := l, l
	for i := l; i <= r; i++ {
		if nums[i] < pivot {
			nums[lt], nums[i] = nums[i], nums[lt]
			if lt < eq {
				nums[eq], nums[i] = nums[i], nums[eq]
			}
			lt++
			eq++
		} else if nums[i] == pivot {
			nums[eq], nums[i] = nums[i], nums[eq]
			eq++
		}
	}
	return lt, eq
}
