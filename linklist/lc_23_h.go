package linklist

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return mergeSort(lists, 0, len(lists)-1)
}

func mergeSort(lists []*ListNode, l, r int) *ListNode {
	if l == r {
		return lists[l]
	}
	mid := l + (r-l)/2
	left := mergeSort(lists, l, mid)
	right := mergeSort(lists, mid+1, r)
	return merge(left, right)
}

func merge(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	p := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next

	}
	if l1 != nil {
		p.Next = l1
	}
	if l2 != nil {
		p.Next = l2
	}
	return dummy.Next
}
