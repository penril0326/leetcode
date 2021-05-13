package merge_kth_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if 0 == len(lists) {
		return nil
	}

	n := len(lists)
	for n > 1 {
		k := (n + 1) / 2
		for i := 0; i < (n / 2); i++ {
			lists[i] = mergeTwoList(lists[i], lists[i+k])
		}

		n = k
	}

	return lists[0]
}

func mergeTwoList(l1, l2 *ListNode) *ListNode {
	head := &ListNode{
		Val: -1,
	}

	current := head
	for (nil != l1) && (nil != l2) {
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}

		current = current.Next
	}

	if nil != l1 {
		current.Next = l1
	}

	if nil != l2 {
		current.Next = l2
	}

	return head.Next
}
