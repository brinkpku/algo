package binarytree

//https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/

import "github.com/brinkpku/algo/utils"

// maxDepth 分解子问题解法
func lc104maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := lc104maxDepth(root.Left)
	rightDepth := lc104maxDepth(root.Right)
	//postorder
	return utils.MaxInt(leftDepth, rightDepth) + 1
}

// maxDepth2 遍历解法
func lc104maxDepth2(root *TreeNode) int {
	depth := 0
	res := 0
	lc104traverse(root, &depth, &res)
	return res
}

func lc104traverse(root *TreeNode, depth, storeRes *int) {
	if root == nil {
		return
	}
	//preorder
	*depth++
	lc104traverse(root.Left, depth, storeRes)
	lc104traverse(root.Right, depth, storeRes)
	//postorder
	if root.Left == nil && root.Right == nil { //可以只在叶节点更新
		*storeRes = utils.MaxInt(*storeRes, *depth)
	}
	*depth--
}
