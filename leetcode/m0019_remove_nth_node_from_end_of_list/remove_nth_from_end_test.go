package remove_nth

import (
	"practice/data_structure/node"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	head := node.ListNode{
		Val: 1,
		Next: &node.ListNode{
			Val: 2,
			Next: &node.ListNode{
				Val: 3,
				Next: &node.ListNode{
					Val: 4,
					Next: &node.ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	removeNthFromEnd(&head, 2)

	current := &head
	for current.Next != nil {
		t.Log(current.Val)
		current = current.Next
	}
	t.Log(current.Val)
}
