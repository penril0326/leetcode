package middleoflinkedlist

import "practice/data_structure/node"

// Time: O(N)
// Space: O(1)
func middleNode(head *node.ListNode) *node.ListNode {
	slow, fast := head, head
	for (fast != nil) && (fast.Next != nil) {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
