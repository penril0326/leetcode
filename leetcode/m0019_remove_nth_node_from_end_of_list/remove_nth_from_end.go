package remove_nth

import "practice/data_structure/node"

// Time: O(N)
// Space: O(1)
func removeNthFromEnd(head *node.ListNode, n int) *node.ListNode {
	if head == nil {
		return nil
	}

	fast := head
	for n > 0 {
		fast = fast.Next
		n--
	}

	dummy := new(node.ListNode)
	dummy.Next = head
	slow := dummy
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	if slow.Next != nil {
		slow.Next = slow.Next.Next
	}

	return dummy.Next
}
