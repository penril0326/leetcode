package validparentheses

import "practice/data_structure/stack"

// Time complexity: O(N)
// Space complexity: O(N)
func isValid(s string) bool {
	st := stack.NewStack()
	for _, r := range s {
		top, _ := st.Top().(rune)
		if isPair(top, r) {
			st.Pop()
		} else {
			st.Push(r)
		}
	}

	return st.IsEmpty()
}

func isPair(left, right rune) bool {
	switch left {
	case '(':
		return right == ')'
	case '{':
		return right == '}'
	case '[':
		return right == ']'
	default:
		return false
	}
}
