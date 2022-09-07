package stack

type MyStack []interface{}

func NewStack() MyStack {
	return make([]interface{}, 0)
}

func (s MyStack) Top() interface{} {
	if len(s) == 0 {
		return nil
	}

	return s[len(s)-1]
}

func (s *MyStack) Pop() interface{} {
	if s == nil || len(*s) == 0 {
		return nil
	}

	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return top
}

func (s *MyStack) Push(v interface{}) {
	if s == nil {
		*s = NewStack()
	}

	*s = append(*s, v)
}

func (s MyStack) Len() int {
	return len(s)
}
