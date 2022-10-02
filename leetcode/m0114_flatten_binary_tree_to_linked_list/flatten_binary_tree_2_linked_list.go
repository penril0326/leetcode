package flattenbinarytreetolinkedlist

import "practice/leetcode"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Time complexity: O(n)
// Space complexity: O(n)
type myStack struct {
	stack []*leetcode.TreeNode
}

func constructor() myStack {
	return myStack{
		stack: make([]*leetcode.TreeNode, 0),
	}
}

func (s myStack) top() *leetcode.TreeNode {
	if !s.isEmpty() {
		return s.stack[len(s.stack)-1]
	}

	return nil
}

func (s myStack) isEmpty() bool {
	return len(s.stack) == 0
}

func (s *myStack) push(n *leetcode.TreeNode) {
	if s != nil {
		s.stack = append(s.stack, n)
	}
}

func (s *myStack) pop() *leetcode.TreeNode {
	if (s != nil) && !s.isEmpty() {
		top := s.top()
		s.stack = s.stack[:len(s.stack)-1]
		return top
	}

	return nil
}

func flatten(root *leetcode.TreeNode) {
	if root == nil {
		return
	}

	stack := constructor()
	stack.push(root)
	for !stack.isEmpty() {
		cur := stack.pop()
		if cur.Right != nil {
			stack.push(cur.Right)
		}

		if cur.Left != nil {
			stack.push(cur.Left)
		}

		if cur != root {
			root.Right = cur
			root.Left = nil
			root = root.Right
		}
	}
}

// -----------------------------
// Time complexity: O(n)
// Space complexity: O(n)
func flattenRecursive(root *leetcode.TreeNode) {
	preorder(root)
}

func preorder(node *leetcode.TreeNode) *leetcode.TreeNode {
	if node == nil {
		return nil
	}

	if (node.Right == nil) && (node.Left == nil) {
		return node
	}

	leftTail := preorder(node.Left)
	rightTail := preorder(node.Right)

	if leftTail != nil {
		leftTail.Right = node.Right
		node.Right = node.Left
		node.Left = nil
	}

	if rightTail == nil {
		return leftTail
	}

	return rightTail
}
