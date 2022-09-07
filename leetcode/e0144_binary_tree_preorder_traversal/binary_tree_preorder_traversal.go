package binarytreepreordertraversal

import "practice/leetcode"

type myStack []*leetcode.TreeNode

func (s myStack) isEmpty() bool {
	return len(s) == 0
}

func (s *myStack) Push(node *leetcode.TreeNode) {
	if s == nil {
		*s = make([]*leetcode.TreeNode, 0)
	}

	*s = append(*s, node)
}

func (s *myStack) Pop() *leetcode.TreeNode {
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return top
}

func preorderTraversal(root *leetcode.TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := make([]int, 0)
	stack := myStack{}
	stack.Push(root)
	for stack.isEmpty() == false {
		cur := stack.Pop()
		result = append(result, cur.Val)

		if cur.Right != nil {
			stack.Push(cur.Right)
		}

		if cur.Left != nil {
			stack.Push(cur.Left)
		}
	}

	return result
}

func preorderTraversalRecursive(root *leetcode.TreeNode) []int {
	result := make([]int, 0)
	traversal(root, &result)

	return result
}

// root, left ,right
func traversal(root *leetcode.TreeNode, result *[]int) {
	if root != nil {
		*result = append(*result, root.Val)
		traversal(root.Left, result)
		traversal(root.Right, result)
	}
}
