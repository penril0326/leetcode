package removeduplicatesfromsortedlist

import "practice/leetcode"

func deleteDuplicates(head *leetcode.ListNode) *leetcode.ListNode {
	if head == nil {
		return nil
	} else if head.Next == nil {
		return head
	}

	slow, fast := head, head.Next
	for slow.Next != nil && fast.Next != nil {
		if slow.Val == fast.Val {
			slow.Next = fast.Next
		} else {
			slow = fast
		}

		fast = fast.Next
	}

	if slow != fast {
		if slow.Val == fast.Val {
			slow.Next = nil
		}
	}

	return head
}
