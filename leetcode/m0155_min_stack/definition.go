package minstack

type MinStack interface {
	Constructor() MinStack
	Push(int)
	Pop()
	Top() int
	GetMin() int
}
