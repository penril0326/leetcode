package flattenbinarytreetolinkedlist

import (
	"practice/data_structure/node"
)

// Time complexity: O(n)
// Space complexity: O(1)
func flatten(root *node.TreeNode) {
	if root == nil {
		return
	}

	cur := root
	for cur != nil {
		if cur.Left != nil {
			rightMost := cur.Left
			for rightMost.Right != nil {
				rightMost = rightMost.Right
			}

			rightMost.Right = cur.Right
			cur.Right = cur.Left
			cur.Left = nil
		}

		cur = cur.Right
	}
}

// -----------------------------
// Time complexity: O(n)
// Space complexity: O(n)
type myStack struct {
	stack []*node.TreeNode
}

func constructor() myStack {
	return myStack{
		stack: make([]*node.TreeNode, 0),
	}
}

func (s myStack) top() *node.TreeNode {
	if !s.isEmpty() {
		return s.stack[len(s.stack)-1]
	}

	return nil
}

func (s myStack) isEmpty() bool {
	return len(s.stack) == 0
}

func (s *myStack) push(n *node.TreeNode) {
	if s != nil {
		s.stack = append(s.stack, n)
	}
}

func (s *myStack) pop() *node.TreeNode {
	if (s != nil) && !s.isEmpty() {
		top := s.top()
		s.stack = s.stack[:len(s.stack)-1]
		return top
	}

	return nil
}

func flattenStack(root *node.TreeNode) {
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
func flattenRecursive(root *node.TreeNode) {
	preorder(root)
}

func preorder(node *node.TreeNode) *node.TreeNode {
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
