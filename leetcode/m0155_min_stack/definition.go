package minstack

type MinStack interface {
	Push(int)
	Pop()
	Top() int
	GetMin() int
}
