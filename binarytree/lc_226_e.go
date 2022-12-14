package binarytree

func lc226invertTree(root *Node) *Node {
	lc226invert(root)
	return root
}

// lc226invert εεΊιε
func lc226invert(root *Node) {
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
