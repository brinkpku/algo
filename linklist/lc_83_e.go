package linklist

// 递归删除重复节点
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := head
	for p != nil && p.Next != nil && p.Val == p.Next.Val {
		p = p.Next
	}
	head.Next = deleteDuplicates(p.Next)
	return head
}

// 迭代删除重复节点
func deleteDuplicates2(head *ListNode) *ListNode {
	p := head
	for p != nil && p.Next != nil {
		if p.Val != p.Next.Val {
			p = p.Next
		} else {
			p.Next = p.Next.Next
		}
	}
	return head
}
