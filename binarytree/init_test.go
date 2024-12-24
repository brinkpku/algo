package binarytree

var (
	a  = &TreeNode{Val: "a"}
	b  = &TreeNode{Val: "b"}
	c  = &TreeNode{Val: "c"}
	d  = &TreeNode{Val: "d"}
	n1 = &TreeNode{Val: 1}
	n2 = &TreeNode{Val: 2}
	n3 = &TreeNode{Val: 3}
	/*
			 a
			/ \
		   b   c
		  / \ / \
		 d  1 2  3
	*/
	fullTree, _ = BinaryTree(a, b, c, d, n1, n2, n3)
)
