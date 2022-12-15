package maximumtwinssumofalinkedlist

import (
	"practice/data_structure/node"
	"practice/utility/math"
)

// Time: O(N)
// Space: O(1)
func pairSum(head *node.ListNode) int {
	if head == nil {
		return 0
	}

	slow, fast := head, head.Next
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	half := reverse(slow.Next)
	maxSum := 0
	for half != nil {
		tmp := head.Val + half.Val
		maxSum = math.Max(maxSum, tmp)
		head = head.Next
		half = half.Next
	}

	return maxSum
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
