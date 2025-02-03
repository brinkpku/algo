// Package binarysearch related
package binarysearch

// 标准的二分搜索框架，搜索目标元素的索引，若不存在则返回 -1
func search(nums []int, target int) int {
	left := 0
	// 注意
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			// 注意
			left = mid + 1
		} else if nums[mid] > target {
			// 注意
			right = mid - 1
		}
	}
	// 如果返回left，则为插入点
	return -1
}

// 将给定的数字 nums 搜索 target 的左侧边界
func leftBound(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	// 搜索区间为 [left, right]
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			// 搜索区间变为 [mid+1, right]
			left = mid + 1
		} else if nums[mid] > target {
			// 搜索区间变为 [left, mid-1]
			right = mid - 1
		} else if nums[mid] == target {
			// 收缩右侧边界
			right = mid - 1
		}
	}
	// 判断 target 是否存在于 nums 中
	if left < 0 || left >= len(nums) {
		return -1
	}
	// 如果越界，target 肯定不存在，返回 -1
	if nums[left] == target {
		// 判断一下 nums[left] 是不是 target
		return left
	}
	return -1
}

func rightBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			// 这里改成收缩左侧边界即可
			left = mid + 1
		}
	}
	// 最后改成返回 left - 1
	if right < 0 || right >= len(nums) {
		return -1
	}
	if nums[right] == target {
		return right
	}
	return -1
}
