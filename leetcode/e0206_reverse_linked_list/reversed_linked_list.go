package reverselinkedlist

import "practice/leetcode"

func reverseList(head *leetcode.ListNode) *leetcode.ListNode {
	var previous *leetcode.ListNode = nil
	for head != nil {
		current := head
		head = head.Next
		current.Next = previous
		previous = current
	}

	return previous
}

func reverseListRecursive(head *leetcode.ListNode) *leetcode.ListNode {
	if (head == nil) || (head.Next == nil) {
		return head
	}

	newHead := reverseListRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil

	return newHead
}
