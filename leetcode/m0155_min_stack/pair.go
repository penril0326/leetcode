package minstack

type pairStack struct {
	stack [][2]int
}

func ConstructorPairStack() pairStack {
	return pairStack{
		stack: make([][2]int, 0),
	}
}

func (s *pairStack) Push(val int) {
	if s != nil {
		min := s.stack[len(s.stack)-1][1]
		if val <= min {
			min = val
		}
		pair := [2]int{val, min}
		s.stack = append(s.stack, pair)
	}
}

func (s *pairStack) Pop() {
	if (s != nil) && (len(s.stack) > 0) {
		s.stack = s.stack[:len(s.stack)-1]
	}
}

func (s *pairStack) Top() int {
	if (s != nil) && (len(s.stack) > 0) {
		return s.stack[len(s.stack)-1][0]
	}

	return -1
}

func (s *pairStack) GetMin() int {
	if (s != nil) && (len(s.stack) > 0) {
		return s.stack[len(s.stack)-1][1]
	}

	return -1
}
