package datastructure

type myStack struct {
	stack []interface{}
}

func NewStack() myStack {
	return myStack{
		stack: make([]interface{}, 0),
	}
}

func (s myStack) IsEmpty() bool {
	return len(s.stack) == 0
}

func (s myStack) Top() interface{} {
	if !s.IsEmpty() {
		return s.stack[len(s.stack)-1]
	}

	return nil
}

func (s *myStack) Push(node interface{}) {
	if s != nil {
		s.stack = append(s.stack, node)
	}
}

func (s *myStack) Pop() interface{} {
	if (s != nil) && !s.IsEmpty() {
		top := s.Top()
		s.stack = s.stack[:len(s.stack)-1]
		return top
	}

	return nil
}
