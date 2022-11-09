package palindromelinkedlist

import (
	"practice/data_structure/node"
)

// Time complexity: O(n)
// Space complexity: O(n)
type myStack struct {
	stack []*node.ListNode
}

func constructor() myStack {
	return myStack{
		stack: make([]*node.ListNode, 0),
	}
}

func (s myStack) isEmpty() bool {
	return len(s.stack) == 0
}

func (s myStack) top() *node.ListNode {
	if len(s.stack) > 0 {
		return s.stack[len(s.stack)-1]
	}

	return nil
}

func (s *myStack) push(node *node.ListNode) {
	if node != nil {
		s.stack = append(s.stack, node)
	}
}

func (s *myStack) pop() *node.ListNode {
	if len(s.stack) > 0 {
		top := s.top()
		s.stack = s.stack[:len(s.stack)-1]
		return top
	}

	return nil
}

func isPalindrome(head *node.ListNode) bool {
	if head == nil {
		return true
	}

	slow, fast := head, head.Next
	stack := constructor()
	for (fast.Next != nil) && (fast.Next.Next != nil) {
		stack.push(slow)
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast.Next != nil {
		stack.push(slow)
	}

	slow = slow.Next

	for slow != nil {
		top := stack.pop()
		if top.Val != slow.Val {
			return false
		}

		slow = slow.Next
	}

	return true
}

// -----------------------------
// Time complexity: O(n)
// Space complexity: O(1)
// This approach should be medium level
func isPalindromeSplit(head *node.ListNode) bool {
	if head == nil {
		return true
	}

	firstHalfEnd := findHalf(head)
	secondHalfStart := reverseLinkedList(firstHalfEnd.Next)

	p1, p2 := head, secondHalfStart
	for p2 != nil {
		if p1.Val != p2.Val {
			return false
		}

		p1 = p1.Next
		p2 = p2.Next
	}

	firstHalfEnd.Next = reverseLinkedList(secondHalfStart)
	return true
}

func reverseLinkedList(head *node.ListNode) *node.ListNode {
	var prev *node.ListNode
	for head != nil {
		cur := head
		head = head.Next
		cur.Next = prev
		prev = cur
	}

	return prev
}

func findHalf(head *node.ListNode) *node.ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}

// -----------------------------
// Time complexity: O(n)
// Space complexity: O(n)
var front *node.ListNode

func isPalindromeRecursive(head *node.ListNode) bool {
	front = head
	return checkPalindrome(head)
}

func checkPalindrome(head *node.ListNode) bool {
	// treat nil is valid palindrome
	if head == nil {
		return true
	}

	// let head mofe to the end node
	if !checkPalindrome(head.Next) {
		return false
	}

	// if head.Next not nil, compare to front node
	if front.Val == head.Val {
		// forward front to next, so far the linked list is valid palindrome
		front = front.Next
		return true
	}

	return false
}
