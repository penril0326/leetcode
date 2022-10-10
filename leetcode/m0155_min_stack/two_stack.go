package minstack

type twoStack struct {
	stack    []int
	minStack []int
}

func ConstructorTwoStack() twoStack {
	return twoStack{
		stack:    make([]int, 0),
		minStack: make([]int, 0),
	}
}

func (s *twoStack) Push(val int) {
	if s != nil {
		s.stack = append(s.stack, val)
		if (len(s.minStack) == 0) || (val <= s.minStack[len(s.minStack)-1]) {
			s.minStack = append(s.minStack, val)
		}
	}
}

func (s *twoStack) Pop() {
	if (s != nil) && (len(s.stack) > 0) && (len(s.minStack) > 0) {
		if s.stack[len(s.stack)-1] == s.minStack[len(s.minStack)-1] {
			s.minStack = s.minStack[:len(s.minStack)-1]
		}

		s.stack = s.stack[:len(s.stack)-1]
	}
}

func (s *twoStack) Top() int {
	if (s != nil) && (len(s.stack) > 0) {
		return s.stack[len(s.stack)-1]
	}

	return -1
}

func (s *twoStack) GetMin() int {
	if (s != nil) && (len(s.minStack) > 0) {
		return s.minStack[len(s.minStack)-1]
	}

	return -1
}
