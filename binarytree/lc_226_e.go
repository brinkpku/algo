package binarytree

func lc226invertTree(root *TreeNode) *TreeNode {
	lc226invert(root)
	return root
}

// lc226invert 前序遍历
func lc226invert(root *TreeNode) {
	if root == nil {
		return
	}
	// preorder
	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp
	lc226invert(root.Left)
	lc226invert(root.Right)
}
