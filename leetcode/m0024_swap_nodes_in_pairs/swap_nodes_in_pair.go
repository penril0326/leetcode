package swap_nodes_in_pair

import "practice/data_structure/node"

// Iterative
// Time: O(N)
// Space: O(1)
func swapPairs(head *node.ListNode) *node.ListNode {
	if (head == nil) || (head.Next == nil) {
		return head
	}

	first := new(node.ListNode) // first
	first.Next = head           // first -> 1 -> 2 -> 3 -> 4
	prev := first               // first(prev) -> 1 -> 2 -> 3 -> 4
	for (head != nil) && (head.Next != nil) {
		keepNext := head.Next.Next // first(prev) -> 1(head) -> 2 -> 3(keepNext) -> 4
		head.Next.Next = head      // first(prev) -> 1(head) <-> 2  3(keepNext) -> 4
		prev.Next = head.Next      // first(prev) -> 2 <-> 1(head)  3(keepNext) -> 4 (prev = first)
		head.Next = keepNext       // first -> 2 -> 1(head) -> 3(keepNext) -> 4

		prev = head     // first -> 2 -> 1(head, prev) -> 3(keepNext) -> 4
		head = keepNext // first -> 2 -> 1(prev) -> 3(head, keepNext) -> 4
	}

	return first.Next
}

// Recursive
// Time: O(N)
// Space: O(N)
func swapPairs2(head *node.ListNode) *node.ListNode {
	if (head == nil) || (head.Next == nil) {
		return head
	}

	first := head
	second := head.Next
	// first -> second -> (tail)

	first.Next = swapPairs(second.Next)
	// first -> (swap)
	// second -> ???

	second.Next = first
	// second -> first -> (swap)

	return second
}
