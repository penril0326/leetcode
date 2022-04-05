package removelinkedlistelements

import "practice/leetcode"

func removeElements(head *leetcode.ListNode, val int) *leetcode.ListNode {
	previous := &leetcode.ListNode{
		Next: head,
	}

	current := previous
	for current.Next != nil {
		if current.Next.Val == val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return previous.Next
}

func removeElementsRecursive(head *leetcode.ListNode, val int) *leetcode.ListNode {
	if head == nil {
		return nil
	}

	head.Next = removeElementsRecursive(head.Next, val)

	if head.Val == val {
		return head.Next
	} else {
		return head
	}
}
