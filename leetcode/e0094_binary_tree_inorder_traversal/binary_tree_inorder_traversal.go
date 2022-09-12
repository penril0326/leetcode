package binarytreeinordertraversal

import "practice/leetcode"

type myStack []*leetcode.TreeNode

func (s *myStack) Push(v *leetcode.TreeNode) {
	if len(*s) == 0 {
		*s = make([]*leetcode.TreeNode, 0)
	}

	*s = append(*s, v)
}

func (s *myStack) Pop() *leetcode.TreeNode {
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return top
}

func (s myStack) isEmpty() bool {
	return len(s) == 0
}

// stack
func inorderTraversal(root *leetcode.TreeNode) []int {
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
func inorderTraversalRecursive(root *leetcode.TreeNode) []int {
	result := make([]int, 0)
	traversal(root, &result)

	return result
}

func traversal(root *leetcode.TreeNode, result *[]int) {
	if root != nil {
		traversal(root.Left, result)
		*result = append(*result, root.Val)
		traversal(root.Right, result)
	}
}
