package palindromelinkedlist

import "practice/leetcode"

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
