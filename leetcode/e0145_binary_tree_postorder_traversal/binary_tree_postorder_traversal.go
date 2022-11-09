package binarytreepostordertraversal

import "practice/data_structure/node"

// Stack
func postorderTraversal(root *node.TreeNode) []int {
	if root == nil {
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

type myStack []*node.TreeNode

func (s myStack) isEmpty() bool {
	return len(s) == 0
}

func (s *myStack) pop() *node.TreeNode {
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return top
}

func (s *myStack) push(n *node.TreeNode) {
	if s == nil {
		*s = make([]*node.TreeNode, 0)
	}

	*s = append(*s, n)
}

func (s myStack) top() *node.TreeNode {
	return s[len(s)-1]
}

// Recursive
func postorderTraversalRecursive(root *node.TreeNode) []int {
	result := make([]int, 0)
	traversal(root, &result)

	return result
}

func traversal(root *node.TreeNode, result *[]int) {
	if root != nil {
		traversal(root.Left, result)
		traversal(root.Right, result)
		*result = append(*result, root.Val)
	}
}
