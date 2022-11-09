package linkedlistcycle

import "practice/leetcode"

func hasCycle(head *leetcode.ListNode) bool {
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
