package binarytree

//https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/

import "github.com/brinkpku/algo/utils"

//maxDepth 分解子问题解法
func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	//postorder
	return utils.MaxInt(leftDepth, rightDepth) + 1
}

var (
	depth = 0
	res   = 0
)

// maxDepth2 遍历解法
func maxDepth2(root *Node) int {
	if root == nil {
		return res
	}
	// preorder
	depth++
	maxDepth2(root.Left)
	maxDepth2(root.Right)
	// postorder
	if root.Left == nil && root.Right == nil { //可以只在叶节点更新
		res = utils.MaxInt(res, depth)
	}
	depth--
	return res
}
