package narytreepostordertraversal

import "practice/data_structure/node"

// Time complexity: O(n)
// Space complexity: O(n)
func postorder(root *node.NaryNode) []int {
	result := make([]int, 0)
	traversal(root, &result)

	return result
}

func traversal(root *node.NaryNode, result *[]int) {
	if root != nil {
		for _, child := range root.Children {
			traversal(child, result)
		}

		*result = append(*result, root.Val)
	}
}

// -----------------------------
// Time complexity: O(n)
// Space complexity: O(n)
type myStack struct {
	stack []*node.NaryNode
}

func constructorStack() myStack {
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

func postorderStack(root *node.NaryNode) []int {
	if root == nil {
		return []int{}
	}

	result := make([]int, 0)
	stack := constructorStack()
	out := constructorStack()
	stack.push(root)
	for !stack.isEmpty() {
		cur := stack.pop()
		out.push(cur)

		for _, child := range cur.Children {
			stack.push(child)
		}
	}

	for !out.isEmpty() {
		top := out.pop()
		result = append(result, top.Val)
	}

	return result
}
