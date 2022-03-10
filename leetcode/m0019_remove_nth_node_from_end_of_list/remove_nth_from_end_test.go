package remove_nth

import (
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	head := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
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
