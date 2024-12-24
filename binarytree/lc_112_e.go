package binarytree

// 回溯所有路径，将sum传在函数中即为dfs遍历
func hasPathSum(root *TreeNode, targetSum int) bool {
	var traverse func(node *TreeNode)
	sum := 0
	ans := false
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		sum += node.Val
		if node.Left == nil && node.Right == nil && sum == targetSum {
			ans = true
			return
		}
		traverse(node.Left)
		traverse(node.Right)
		sum -= node.Val
	}
	traverse(root)
	return ans
}