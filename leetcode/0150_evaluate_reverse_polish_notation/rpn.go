package rpn

import (
	"container/list"
	"strconv"
)

type IntStack struct {
	list *list.List
}

func newStack() *IntStack {
	return &IntStack{
		list: list.New(),
	}
}

func (s *IntStack) Push(value int) {
	s.list.PushFront(value)
}

func (s *IntStack) Pop() int {
	e := s.list.Front()
	value := 0
	if e != nil {
		value = e.Value.(int)
		s.list.Remove(e)
	}

	return value
}

func evalRPN(tokens []string) int {
	stack := newStack()

	for _, token := range tokens {
		value, err := strconv.Atoi(token)
		if err == nil {
			stack.Push(value)
		} else {
			a := stack.Pop()
			b := stack.Pop()
			c := cal(b, a, token)
			stack.Push(c)
		}
	}

	return stack.Pop()
}

func cal(a, b int, operator string) int {
	result := 0
	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	}

	return result
}
