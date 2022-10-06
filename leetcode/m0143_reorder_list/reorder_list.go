package reorderlist

import "practice/leetcode"

// Time complexity: O(n)
// Space complexity: O(n)
type myStack struct {
	stack []*leetcode.ListNode
}

func newStack() myStack {
	return myStack{
		stack: make([]*leetcode.ListNode, 0),
	}
}

func (s myStack) isEmpty() bool {
	return len(s.stack) == 0
}

func (s myStack) top() *leetcode.ListNode {
	if !s.isEmpty() {
		return s.stack[len(s.stack)-1]
	}

	return nil
}

func (s *myStack) push(node *leetcode.ListNode) {
	if s != nil {
		s.stack = append(s.stack, node)
	}
}

func (s *myStack) pop() *leetcode.ListNode {
	if (s != nil) && !s.isEmpty() {
		top := s.top()
		s.stack = s.stack[:len(s.stack)-1]
		return top
	}

	return nil
}

func reorderList(head *leetcode.ListNode) {
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
