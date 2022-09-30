package binarytree

// https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/

func lc114flatten(root *Node) {
	emptyNode := &Node{}
	lc114traverse(root, emptyNode)
	root = emptyNode.Right
}

func lc114traverse(root, res *Node) {
	if root == nil {
		return
	}
	//preorder
	res.Right = &Node{Val: root.Val}
	res = res.Right
	lc114traverse(root.Left, res)
	lc114traverse(root.Right, res)
}
