package reorderlist

import (
	"practice/data_structure/node"
)

// Time complexity: O(n)
// Space complexity: O(n)
type myStack struct {
	stack []*node.ListNode
}

func newStack() myStack {
	return myStack{
		stack: make([]*node.ListNode, 0),
	}
}

func (s myStack) isEmpty() bool {
	return len(s.stack) == 0
}

func (s myStack) top() *node.ListNode {
	if !s.isEmpty() {
		return s.stack[len(s.stack)-1]
	}

	return nil
}

func (s *myStack) push(node *node.ListNode) {
	if s != nil {
		s.stack = append(s.stack, node)
	}
}

func (s *myStack) pop() *node.ListNode {
	if (s != nil) && !s.isEmpty() {
		top := s.top()
		s.stack = s.stack[:len(s.stack)-1]
		return top
	}

	return nil
}

func reorderList(head *node.ListNode) {
	if (head == nil) || (head.Next == nil) {
		return
	}

	s, f := head, head
	for (f.Next != nil) && (f.Next.Next != nil) {
		s = s.Next
		f = f.Next.Next
	}

	stack := newStack()
	for s.Next != nil {
		s = s.Next
		stack.push(s)
	}

	for !stack.isEmpty() {
		tmp := head.Next
		head.Next = stack.pop()
		if tmp == head.Next {
			head.Next.Next = nil
		} else {
			head.Next.Next = tmp
		}

		head = head.Next.Next
	}

	if head != nil {
		head.Next = nil
	}
}

// -----------------------------
// Time complexity: O(n)
// Space complexity: O(n)
func reorderListRecursive(head *node.ListNode) {
	findNextInsert(head, head)
}

func findNextInsert(head, curr *node.ListNode) *node.ListNode {
	if curr.Next != nil {
		head = findNextInsert(head, curr.Next)
	}

	if head == nil {
		return nil
	}

	second := head.Next

	if (curr == head) || (curr == second) {
		curr.Next = nil
		return nil
	}

	head.Next = curr
	curr.Next = second
	return second
}

// -----------------------------
// Time complexity: O(n)
// Space complexity: O(1)
func reorderListIterative(head *node.ListNode) {
	// find middle
	slow, fast := head, head
	for (fast != nil) && (fast.Next != nil) {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// reverse second part
	var prev *node.ListNode
	cur := slow
	for cur != nil {
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}

	// merge two linked lists
	first, second := head, prev
	for second.Next != nil {
		tmp := first.Next
		first.Next = second
		first = tmp

		tmp = second.Next
		second.Next = first
		second = tmp
	}
}
