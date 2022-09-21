package backspacestringcompare

// Time complexity: O(m+n), m and n are length of s and t respectively
// Space complexity: O(m+n)
type myStack struct {
	stack []rune
}

func constructor() myStack {
	return myStack{
		stack: make([]rune, 0),
	}
}

func (s myStack) top() rune {
	if len(s.stack) > 0 {
		return s.stack[len(s.stack)-1]
	}

	return 0
}

func (s *myStack) push(r rune) {
	s.stack = append(s.stack, r)
}

func (s *myStack) pop() rune {
	if len(s.stack) > 0 {
		top := s.top()
		s.stack = s.stack[:len(s.stack)-1]

		return top
	}

	return 0
}

func (s myStack) isEmpty() bool {
	return len(s.stack) == 0
}

func backspaceCompare(s string, t string) bool {
	return handleBackspace(s) == handleBackspace(t)
}

func handleBackspace(s string) string {
	stk := constructor()
	for _, r := range s {
		if r == '#' {
			if !stk.isEmpty() {
				stk.pop()
			}
		} else {
			stk.push(r)
		}
	}

	result := ""
	for !stk.isEmpty() {
		r := stk.pop()
		result += string(r)
	}

	return result
}

// -----------------------------
// Time complexity: O(m+n), m and n are length of s and t respectively
// Space complexity: O(m+n)
func backspaceCompareWithoutStack(s string, t string) bool {
	return handleBackspace2(s) == handleBackspace2(t)
}

func handleBackspace2(s string) string {
	newStr := ""
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '#' {
			count++
		} else if count == 0 {
			newStr += string(s[i])
		} else {
			count--
		}
	}

	return newStr
}

// -----------------------------
// Time complexity: O(m+n), m and n are length of s and t respectively
// Space complexity: O(1)

func backspaceCompareTwoPointers(s string, t string) bool {
	iS, skipS := len(s)-1, 0
	iT, skipT := len(t)-1, 0
	for (iS >= 0) || (iT >= 0) {
		iS, skipS = findSubStringIdx(s, iS, skipS)
		iT, skipT = findSubStringIdx(t, iT, skipT)

		if (iS >= 0) && (iT >= 0) && (s[iS] != t[iT]) {
			return false
		}

		if (iS >= 0) != (iT >= 0) {
			return false
		}

		iS--
		iT--
	}

	return true
}

func findSubStringIdx(s string, idx, skip int) (int, int) {
	for idx >= 0 {
		if s[idx] == '#' {
			skip++
			idx--
		} else if skip > 0 {
			skip--
			idx--
		} else {
			break
		}
	}

	return idx, skip
}
