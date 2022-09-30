package binarytree

var (
	a  = &Node{Val: "a"}
	b  = &Node{Val: "b"}
	c  = &Node{Val: "c"}
	d  = &Node{Val: "d"}
	n1 = &Node{Val: 1}
	n2 = &Node{Val: 2}
	n3 = &Node{Val: 3}
	/*
			 a
			/ \
		   b   c
		  / \ / \
		 d  1 2  3
	*/
	fullTree, _ = BinaryTree(a, b, c, d, n1, n2, n3)
)
