package makethestringgreat

// Time complexity: O(n)
// Space complexity: O(n)
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
	if s != nil {
		s.stack = append(s.stack, r)
	}
}

func (s *myStack) pop() rune {
	if (s != nil) && (len(s.stack) > 0) {
		top := s.top()
		s.stack = s.stack[:len(s.stack)-1]
		return top
	}

	return 0
}

func (s myStack) isEmpty() bool {
	return len(s.stack) == 0
}

func makeGood(s string) string {
	if s == "" {
		return ""
	}

	stack := constructor()
	for _, r := range s {
		if !stack.isEmpty() {
			if isSameLetterWithCaseInsenstive(stack.top(), r) {
				stack.pop()
				continue
			}
		}

		stack.push(r)
	}

	result := []rune{}
	for !stack.isEmpty() {
		result = append(result, stack.pop())
	}

	// reverse
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}

func isSameLetterWithCaseInsenstive(r1, r2 rune) bool {
	if ('a' <= r1) && (r1 <= 'z') {
		// check if r2 is r1's uppercase
		if (r2 - 'A' + 'a') == r1 {
			return true
		}
	} else if ('A' <= r1) && (r1 <= 'Z') {
		// check if r2 is r1's lowercase
		if (r2 - 'a' + 'A') == r1 {
			return true
		}
	}

	return false
}

// -----------------------------
// Time complexity: O(n^2)
// Space complexity: O(n^2)
func makeGoodRecursive(s string) string {
	for i := 0; i < len(s)-1; i++ {
		if isSameLetterWithCaseInsenstive(rune(s[i]), rune(s[i+1])) {
			return makeGood(s[0:i] + s[i+2:])
		}
	}

	return s
}

// -----------------------------
// Time complexity: O(n^2)
// Space complexity: O(n)
func makeGoodBruteForce(s string) string {
	for len(s) > 1 {
		isFind := false
		for i := 0; i < len(s)-1; i++ {
			if isSameLetterWithCaseInsenstive(rune(s[i]), rune(s[i+1])) {
				s = s[0:i] + s[i+2:]
				isFind = true
				break
			}
		}

		if !isFind {
			break
		}
	}

	return s
}
