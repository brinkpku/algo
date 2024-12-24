package binarytree

func countPath(root *TreeNode, targetSum int) int {
	var traverse func(*TreeNode, int)
	ans := 0
	traverse = func(node *TreeNode, left int) {
		if node == nil {
			return
		}
		left -= node.Val
		if left == 0 {
			ans++
		}
		traverse(node.Left, left)
		traverse(node.Right, left)
	}
	traverse(root, targetSum)
	return ans
}

// 分治+dfs遍历
func pathSum1(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	return pathSum(root.Left, targetSum) + countPath(root, targetSum) + pathSum(root.Right, targetSum)
}