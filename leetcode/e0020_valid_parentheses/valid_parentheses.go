package validparentheses

import "practice/leetcode"

// Time complexity: O(N)
// Space complexity: O(N)
func isValid(s string) bool {
	stack := leetcode.NewStack()
	for _, r := range s {
		top, _ := stack.Top().(rune)
		if isPair(top, r) {
			stack.Pop()
		} else {
			stack.Push(r)
		}
	}

	return stack.IsEmpty()
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
