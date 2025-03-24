package linklist

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	p := dummy
	step := right - left + 1
	var p0 *ListNode
	for ; left > 0; left-- {
		p0 = p
		p = p.Next
	}
	p1 := p
	var prev *ListNode
	for ; p != nil && step > 0; step-- {
		next := p.Next
		p.Next = prev
		prev = p
		p = next
	}
	p0.Next = prev
	p1.Next = p
	return dummy.Next
}
