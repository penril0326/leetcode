package binarytreepostordertraversal

import "practice/leetcode"

// Stack
func postorderTraversal(root *leetcode.TreeNode) []int {
	if root != nil {
		return []int{}
	}

	result := make([]int, 0)
	stack := myStack{}
	stack.push(root)

	out := myStack{}

	for !stack.isEmpty() {
		node := stack.pop()
		out.push(node)

		if node.Left != nil {
			stack.push(node.Left)
		}

		if node.Right != nil {
			stack.push(node.Right)
		}
	}

	for !out.isEmpty() {
		result = append(result, out.pop().Val)
	}

	return result
}

type myStack []*leetcode.TreeNode

func (s myStack) isEmpty() bool {
	return len(s) == 0
}

func (s *myStack) pop() *leetcode.TreeNode {
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return top
}

func (s *myStack) push(n *leetcode.TreeNode) {
	if s == nil {
		*s = make([]*leetcode.TreeNode, 0)
	}

	*s = append(*s, n)
}

func (s myStack) top() *leetcode.TreeNode {
	return s[len(s)-1]
}

// Recursive
func postorderTraversalRecursive(root *leetcode.TreeNode) []int {
	result := make([]int, 0)
	traversal(root, &result)

	return result
}

func traversal(root *leetcode.TreeNode, result *[]int) {
	if root != nil {
		traversal(root.Left, result)
		traversal(root.Right, result)
		*result = append(*result, root.Val)
	}
}
