package removeduplicatesfromsortedlist

import "practice/data_structure/node"

// Time: O(N)
// Space: O(1)
func deleteDuplicates(head *node.ListNode) *node.ListNode {
	current := head
	for (current != nil) && (current.Next != nil) {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return head
}
