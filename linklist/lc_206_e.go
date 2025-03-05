package linklist

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	rest := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return rest
}
