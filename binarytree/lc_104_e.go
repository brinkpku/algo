package binarytree

//https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/

import "github.com/brinkpku/algo/utils"

//maxDepth 分解子问题解法
func lc104maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	leftDepth := lc104maxDepth(root.Left)
	rightDepth := lc104maxDepth(root.Right)
	//postorder
	return utils.MaxInt(leftDepth, rightDepth) + 1
}

var (
	lc104depth = 0
	lc104      = 0
)

// maxDepth2 遍历解法
func lc104maxDepth2(root *Node) int {
	if root == nil {
		return lc104
	}
	// preorder
	lc104depth++
	lc104maxDepth2(root.Left)
	lc104maxDepth2(root.Right)
	// postorder
	if root.Left == nil && root.Right == nil { //可以只在叶节点更新
		lc104 = utils.MaxInt(lc104, lc104depth)
	}
	lc104depth--
	return lc104
}
