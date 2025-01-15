// Package binarytree binary tree common struct and functions
package binarytree

import (
	"fmt"
	"strings"
)

// TreeNode binary tree node
type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

type Position int

const (
	//root node
	RootPos Position = iota
	//left sub tree
	LeftPos
	//right sub tree
	RightPos
)

func (n *TreeNode) String() string {
	return fmt.Sprintf("%v", n.Val)
}

// VerticalPrint print binary tree vertically
func (n *TreeNode) VerticalPrint() {
	verticalPrint(n, RootPos, []string{})
}

// verticalPrint
func verticalPrint(node *TreeNode, pos Position, indnt []string) {
	if node == nil {
		return
	}
	//preorder
	var (
		rindnt []string
		lindnt []string
	)
	if pos == RootPos {
		rindnt = append(indnt, "    ")
		lindnt = append(indnt, "    ")
	} else if pos == RightPos {
		rindnt = append(indnt, "    ")
		lindnt = append(indnt, "|   ")
	} else {
		rindnt = append(indnt, "|   ")
		lindnt = append(indnt, "    ")
	}
	verticalPrint(node.Right, RightPos, rindnt)
	//midorder
	indnt = append(indnt, "|->")
	fmt.Println(strings.Join(indnt, ""), node.String())
	verticalPrint(node.Left, LeftPos, lindnt)
}

/*
dfs遍历框架，前中后序是遍历二叉树过程中处理每一个节点的三个特殊时间点：
前序位置的代码在刚刚进入一个二叉树节点的时候执行；
后序位置的代码在将要离开一个二叉树节点的时候执行；
中序位置的代码在一个二叉树节点左子树都遍历完，即将开始遍历右子树的时候执行。
*/
func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	// preorder code
	traverse(root.Left)
	// midorder code
	traverse(root.Right)
	// postorder code
}

/*
bfs遍历框架，利用队列，并记录层数。
*/
func levelOrderTraverse(root *TreeNode) {
	if root == nil {
		return
	}
	q := []*TreeNode{root}
	// 记录当前遍历到的层数（根节点视为第 1 层）
	depth := 1
	for len(q) > 0 {
		// 获取当前队列长度
		sz := len(q)
		for i := 0; i < sz; i++ {
			// 弹出队列头
			cur := q[0]
			q = q[1:]
			// 访问 cur 节点，同时知道它所在的层数
			fmt.Printf("depth = %d, val = %d\n", depth, cur.Val)
			// 把 cur 的左右子节点加入队列
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		depth++
	}
}

// BinaryTree 贪心方式构造二叉树
func BinaryTree(nodes ...*TreeNode) (root *TreeNode, err error) {
	if len(nodes) == 0 {
		return
	}
	if nodes[0] == nil {
		err = ErrInvalidBinaryTreeNodes
		return
	}
	queue := []*TreeNode{nodes[0]}
	root = nodes[0]
	for i := 1; i < len(nodes); i++ {
		if len(queue) != 0 {
			parent := queue[0]
			queue = queue[1:]
			if nodes[i] != nil {
				parent.Left = nodes[i]
				queue = append(queue, nodes[i])
			}
			i++
			if i < len(nodes) && nodes[i] != nil {
				parent.Right = nodes[i]
				queue = append(queue, nodes[i])
			}
		} else if i < len(nodes) {
			root = nil
			err = ErrInvalidBinaryTreeNodes
			return
		}
	}
	return
}

// cleanNodesConnection clean tree node connection
func cleanNodesConnection(nodes ...*TreeNode) {
	for _, node := range nodes {
		if node != nil {
			node.Left = nil
			node.Right = nil
		}
	}
}
