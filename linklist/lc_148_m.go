package linklist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func sortList(head *ListNode) *ListNode {
	mid := findmid(head)
	if head == mid {
		return head
	}
	left := sortList(head)
	right := sortList(mid)
	dummy := &ListNode{}
	p := dummy
	for left != nil && right != nil {
		if left.Val < right.Val {
			dummy.Next = left
			left = left.Next
		} else {
			dummy.Next = right
			right = right.Next
		}
		dummy = dummy.Next
	}
	if left != nil {
		dummy.Next = left
	}
	if right != nil {
		dummy.Next = right
	}
	return p.Next
}

func findmid(node *ListNode) *ListNode {
	slow, fast := node, node
	prev := slow
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	if prev != nil {
		prev.Next = nil
	}
	return slow
}