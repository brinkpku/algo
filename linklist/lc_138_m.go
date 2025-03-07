package linklist

// Node Definition for a Node.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	p := head
	hash := map[*Node]*Node{}
	dummy := &Node{}
	newNode := dummy
	for p != nil {
		newNode.Next = &Node{
			Val: p.Val,
		}
		hash[p] = newNode.Next
		p = p.Next
		newNode = newNode.Next
	}
	p = head
	newNode = dummy.Next
	for p != nil {
		newNode.Random = hash[p.Random]
		p = p.Next
		newNode = newNode.Next
	}
	return dummy.Next
}
