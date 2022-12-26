package deletethemiddlenodeofalinkedlist

import "practice/data_structure/node"

// Time: O(N)
// Space: O(1)
func deleteMiddle(head *node.ListNode) *node.ListNode {
	if (head == nil) || (head.Next == nil) {
		return nil
	}

	slow, fast := head, head
	dummy := new(node.ListNode)
	dummy.Next = head
	prev := dummy
	for (fast != nil) && (fast.Next != nil) {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	prev.Next = slow.Next

	return dummy.Next
}
