package binarytree

import "math"

func maxPathSum(root *TreeNode) int {
	var maxGain func(node *TreeNode) int
	ans := math.MinInt
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := maxGain(node.Left)
		r := maxGain(node.Right)
		ans = max(ans, node.Val+l+r)
		return max(node.Val+l, node.Val+r, 0)
	}
	maxGain(root)
	return ans
}
