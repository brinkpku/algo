package binarytree

// 回溯所有路径，将sum传在函数中即为dfs遍历
func pathSum(root *TreeNode, targetSum int) [][]int {
	ans := [][]int{}
	temp := []int{}
	sum := 0
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		sum += node.Val
		temp = append(temp, node.Val)
		if sum == targetSum && node.Left == nil && node.Right == nil {
			pathCopy := make([]int, len(temp))
            copy(pathCopy, temp)
			ans = append(ans, pathCopy)
		}
        traverse(node.Left)
        traverse(node.Right)
        temp = temp[:len(temp)-1]
        sum -= node.Val
	}
	traverse(root)
	return ans
}