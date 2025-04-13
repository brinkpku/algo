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
