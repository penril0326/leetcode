package sort_list

import "practice/data_structure/node"

func sortList(head *node.ListNode) *node.ListNode {
	if (head == nil) || (head.Next == nil) {
		return head
	}

	slow, fast, previous := head, head, head
	for (fast != nil) && (fast.Next != nil) {
		previous = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	previous.Next = nil
	return merge(sortList(head), sortList(slow))
}

func merge(left, right *node.ListNode) *node.ListNode {
	head := &node.ListNode{
		Val: -1,
	}

	current := head
	for (left != nil) && (right != nil) {
		if left.Val < right.Val {
			current.Next = left
			left = left.Next
		} else {
			current.Next = right
			right = right.Next
		}

		current = current.Next
	}

	if left != nil {
		current.Next = left
	}

	if right != nil {
		current.Next = right
	}

	return head.Next
}
