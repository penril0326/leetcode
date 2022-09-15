package palindromelinkedlist

import "practice/leetcode"

// Time complexity: O(n)
// Space complexity: O(1)
func isPalindromeSplit(head *leetcode.ListNode) bool {
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

func reverseLinkedList(head *leetcode.ListNode) *leetcode.ListNode {
	var prev *leetcode.ListNode
	for head != nil {
		cur := head
		head = head.Next
		cur.Next = prev
		prev = cur
	}

	return prev
}

func findHalf(head *leetcode.ListNode) *leetcode.ListNode {
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
var front *leetcode.ListNode

func isPalindromeRecursive(head *leetcode.ListNode) bool {
	front = head
	return checkPalindrome(head)
}

func checkPalindrome(head *leetcode.ListNode) bool {
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
