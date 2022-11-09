package swap_nodes_in_pair

import "practice/data_structure/node"

func swapPairs(head *node.ListNode) *node.ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	p := head
	for ; head.Next != nil; head = head.Next {
		swap(head, head.Next)
		if head == nil {
			break
		}
	}

	return p
}

func swap(n1, n2 *node.ListNode) {
	n1.Next = n2.Next
	n2.Next = n1
}
