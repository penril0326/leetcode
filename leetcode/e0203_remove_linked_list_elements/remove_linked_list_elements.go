package removelinkedlistelements

import "practice/data_structure/node"

func removeElements(head *node.ListNode, val int) *node.ListNode {
	previous := &node.ListNode{
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

func removeElementsRecursive(head *node.ListNode, val int) *node.ListNode {
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
