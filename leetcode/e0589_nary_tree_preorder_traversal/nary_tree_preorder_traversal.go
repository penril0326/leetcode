package narytreepreordertraversal

import "practice/data_structure/node"

// Time complexity: O(n)
// Space complexity: O(n)
func preorder(root *node.NaryNode) []int {
	result := make([]int, 0)
	traversal(root, &result)
	return result
}

func traversal(root *node.NaryNode, result *[]int) {
	if root != nil {
		*result = append(*result, root.Val)
		for _, child := range root.Children {
			traversal(child, result)
		}
	}
}

// -----------------------------
// Time complexity: O(n)
// Space complexity: O(n)
type myStack struct {
	stack []*node.NaryNode
}

func stackConstructor() myStack {
	return myStack{
		stack: make([]*node.NaryNode, 0),
	}
}

func (s myStack) top() *node.NaryNode {
	if len(s.stack) > 0 {
		return s.stack[len(s.stack)-1]
	}

	return nil
}

func (s *myStack) push(node *node.NaryNode) {
	s.stack = append(s.stack, node)
}

func (s *myStack) pop() *node.NaryNode {
	if len(s.stack) > 0 {
		top := s.top()
		s.stack = s.stack[:len(s.stack)-1]

		return top
	}

	return nil
}

func (s myStack) isEmpty() bool {
	return len(s.stack) == 0
}

func preorderStack(root *node.NaryNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	stack := stackConstructor()
	stack.push(root)
	for !stack.isEmpty() {
		cur := stack.pop()
		result = append(result, cur.Val)

		for i := len(cur.Children) - 1; i >= 0; i-- {
			stack.push(cur.Children[i])
		}
	}

	return result
}
