package binarytree

// https://leetcode.cn/problems/binary-tree-preorder-traversal/submissions/

func lc144preorderTraversal(root *TreeNode) []interface{} {
	res := make([]interface{}, 0)
	lc144traverse(root, &res)
	return res
}

func lc144traverse(root *TreeNode, res *[]interface{}) {
	if root == nil {
		return
	}
	// pre order
	*res = append(*res, root.Val)
	lc144traverse(root.Left, res)
	lc144traverse(root.Right, res)
}

// todo 使用栈迭代解法
