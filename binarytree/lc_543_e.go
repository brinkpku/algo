package binarytree

// https://leetcode.cn/problems/diameter-of-binary-tree/

func lc543diameterOfBinaryTree(root *Node) int {
	if root == nil {
		return 0
	}
	maxDiameter := 0
	lc543maxDepth(root, &maxDiameter)
	return maxDiameter
}

func lc543maxDepth(root *Node, diam *int) int {
	if root == nil {
		return 0
	}
	leftDepth := lc543maxDepth(root.Left, diam)
	rightDepth := lc543maxDepth(root.Right, diam)
	*diam = lc543max(*diam, leftDepth+rightDepth)
	return lc543max(leftDepth, rightDepth) + 1
}

func lc543max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
