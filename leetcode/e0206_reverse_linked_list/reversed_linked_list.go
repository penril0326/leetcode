package reverselinkedlist

import "practice/data_structure/node"

func reverseList(head *node.ListNode) *node.ListNode {
	var previous *node.ListNode = nil
	for head != nil {
		current := head
		head = head.Next
		current.Next = previous
		previous = current
	}

	return previous
}

func reverseListRecursive(head *node.ListNode) *node.ListNode {
	if (head == nil) || (head.Next == nil) {
		return head
	}

	newHead := reverseListRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil

	return newHead
}
