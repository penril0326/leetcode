package binarytreeinordertraversal

import "practice/data_structure/node"

type myStack []*node.TreeNode

func (s *myStack) Push(v *node.TreeNode) {
	if len(*s) == 0 {
		*s = make([]*node.TreeNode, 0)
	}

	*s = append(*s, v)
}

func (s *myStack) Pop() *node.TreeNode {
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return top
}

func (s myStack) isEmpty() bool {
	return len(s) == 0
}

// stack
func inorderTraversal(root *node.TreeNode) []int {
	cur := root
	result := []int{}
	stack := myStack{}
	for (cur != nil) || (stack.isEmpty() == false) {
		if cur != nil {
			stack.Push(cur)
			cur = cur.Left
		} else {
			cur = stack.Pop()
			result = append(result, cur.Val)
			cur = cur.Right
		}
	}

	return result
}

// recursive
func inorderTraversalRecursive(root *node.TreeNode) []int {
	result := make([]int, 0)
	traversal(root, &result)

	return result
}

func traversal(root *node.TreeNode, result *[]int) {
	if root != nil {
		traversal(root.Left, result)
		*result = append(*result, root.Val)
		traversal(root.Right, result)
	}
}
