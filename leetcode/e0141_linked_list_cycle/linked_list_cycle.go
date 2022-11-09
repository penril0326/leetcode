package linkedlistcycle

import (
	"practice/data_structure/node"
)

func hasCycle(head *node.ListNode) bool {
	slow, fast := head, head
	for (fast != nil) && (fast.Next != nil) {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}
