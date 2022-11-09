package binarytreepreordertraversal

import "practice/data_structure/node"

type myStack []*node.TreeNode

func (s myStack) isEmpty() bool {
	return len(s) == 0
}

func (s *myStack) Push(n *node.TreeNode) {
	if s == nil {
		*s = make([]*node.TreeNode, 0)
	}

	*s = append(*s, n)
}

func (s *myStack) Pop() *node.TreeNode {
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return top
}

func preorderTraversal(root *node.TreeNode) []int {
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

func preorderTraversalRecursive(root *node.TreeNode) []int {
	result := make([]int, 0)
	traversal(root, &result)

	return result
}

// root, left ,right
func traversal(root *node.TreeNode, result *[]int) {
	if root != nil {
		*result = append(*result, root.Val)
		traversal(root.Left, result)
		traversal(root.Right, result)
	}
}
