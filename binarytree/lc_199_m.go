package binarytree

// bfs 解法
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	q := []*TreeNode{root}
	ans := []int{}
	for len(q) > 0 {
		sz := len(q)
		for ; sz > 0; sz-- {
			n := q[0]
			q = q[1:]
			if sz == 1 {
				ans = append(ans, n.Val)
			}
			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
		}
	}
	return ans
}

// dfs 深度优先搜索根据层数记录最右边的节点
func dfsRightSideView(root *TreeNode) []int {
	var traverse func(*TreeNode, int)
	ans := []int{}
	traverse = func(node *TreeNode, depth int) {
		if node == nil {
			return 
		}
		if len(ans) >= depth {
			ans[depth-1] = node.Val
		} else {
			ans = append(ans, node.Val)
		}
		traverse(node.Left, depth+1)
		traverse(node.Right, depth+1)
	}
	traverse(root, 1)
	return ans
}