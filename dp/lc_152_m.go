package dp

import "math"

// 自底向上，dp table
func maxProduct(nums []int) int {
	dpmax := make([]int, len(nums))
	dpmin := make([]int, len(nums))
	ans := math.MinInt
	for i := range nums {
		if i == 0 {
			dpmax[i], dpmin[i] = nums[i], nums[i]
			ans = max(dpmax[i], ans)
		} else {
			dpmax[i] = max(dpmax[i-1]*nums[i], dpmin[i-1]*nums[i], nums[i])
			dpmin[i] = min(dpmax[i-1]*nums[i], dpmin[i-1]*nums[i], nums[i])
			ans = max(dpmax[i], ans)
		}
	}
	return ans
}

// 空间优化
func maxProduct2(nums []int) int {
	dpmax := 1
	dpmin := 1
	ans := math.MinInt
	for _, i := range nums {
		dpmax, dpmin = max(dpmax*i, dpmin*i, i), min(dpmax*i, dpmin*i, i)
		ans = max(dpmax, ans)
	}
	return ans
}
