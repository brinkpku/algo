package linklist

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	p := dummy
	for p != nil && p.Next != nil && p.Next.Next != nil {
		first := p.Next
		second := p.Next.Next
		p.Next = second
		last := second.Next
		second.Next = first
		first.Next = last
		p = first
	}
	return dummy.Next
}

// 递归
func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	second := head.Next
	last := swapPairs(second.Next)
	second.Next = head
	head.Next = last
	return second
}
