package reverselinkedlist2

import "practice/data_structure/node"

// Iterative
// Time: O(N)
// Space: O(1)
func reverseBetween(head *node.ListNode, left int, right int) *node.ListNode {
	if (left == right) || (head == nil) || (head.Next == nil) {
		return head
	}

	var prev, subTail, tail *node.ListNode
	dummy := new(node.ListNode)
	dummy.Next = head

	count := 0
	cur := dummy
	for cur != nil {
		if count+1 == left {
			prev = cur
		} else if count == right {
			subTail = cur
			tail = cur.Next
			break
		}

		cur = cur.Next
		count++
	}

	subTail.Next = nil
	newSubTail := prev.Next
	prev.Next = reverse(prev.Next)

	newSubTail.Next = tail

	return dummy.Next
}

func reverse(head *node.ListNode) *node.ListNode {
	var prev *node.ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}

// Recursive
// Time: O(N)
// Space: O(N)
