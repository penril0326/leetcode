package remove_nth

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	node := &ListNode{
		Val:  0,
		Next: head,
	}

	targetPrevious := node
	current := head

	for n > 0 {
		current = current.Next
		n--
	}

	for current != nil {
		current = current.Next
		targetPrevious = targetPrevious.Next
	}

	if nil != targetPrevious.Next {
		targetPrevious.Next = targetPrevious.Next.Next
	}

	return node.Next
}
