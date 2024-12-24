package binarytree

// https://leetcode.cn/problems/diameter-of-binary-tree/

func diameterOfBinaryTree(root *TreeNode) int {
	ans := 0
	var maxDepth func(n *TreeNode) int
	maxDepth = func(n *TreeNode) int {
		if n == nil {
			return 0
		}
		l := maxDepth(n.Left)
		r := maxDepth(n.Right)
		depth := max(l, r) + 1
		// 求最大深度的时候已经知道了左右子树的深度
		ans = max(ans, l+r)
		return depth
	}
	maxDepth(root)
	return ans
}
