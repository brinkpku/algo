package linklist

func reverseKGroup(head *ListNode, k int) *ListNode {
	p := head
	var last *ListNode
	count := k
	for count = k; count > 0 && p != nil; count-- {
		last = p
		p = p.Next
	}
	if count > 0 {
		return head
	}
	last.Next = nil
	first := reverseList(head)
	rest := reverseKGroup(p, k)
	head.Next = rest
	return first
}
