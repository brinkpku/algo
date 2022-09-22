//Package binarytree binary tree common struct and functions
package binarytree

import "fmt"

// Node binary tree node
type Node struct {
	Left  *Node
	Right *Node
	Val   interface{}
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.Val)
}

func (n *Node) TraversePrint()

/*
遍历框架，前中后序是遍历二叉树过程中处理每一个节点的三个特殊时间点：
前序位置的代码在刚刚进入一个二叉树节点的时候执行；
后序位置的代码在将要离开一个二叉树节点的时候执行；
中序位置的代码在一个二叉树节点左子树都遍历完，即将开始遍历右子树的时候执行。
*/
func traverse(root *Node) {
	if root == nil {
		return
	}
	// preorder code
	traverse(root.Left)
	// midorder code
	traverse(root.Right)
	// postorder code
}

//BinaryTree 贪心方式构造二叉树
func BinaryTree(nodes ...*Node) (root *Node, err error) {
	if len(nodes) == 0 {
		return
	}
	if nodes[0] == nil {
		err = ErrInvalidBinaryTreeNodes
		return
	}
	queue := []*Node{nodes[0]}
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

//cleanNodesConnection clean tree node connection
func cleanNodesConnection(nodes ...*Node) {
	for _, node := range nodes {
		if node != nil {
			node.Left = nil
			node.Right = nil
		}
	}
}
