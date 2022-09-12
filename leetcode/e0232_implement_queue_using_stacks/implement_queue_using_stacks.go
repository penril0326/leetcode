package implementqueueusingstacks

type myStack struct {
	s []int
}

func constructor() myStack {
	return myStack{
		s: make([]int, 0),
	}
}

func (s *myStack) push(x int) {
	s.s = append(s.s, x)
}

func (s *myStack) pop() int {
	if len(s.s) > 0 {
		top := s.s[len(s.s)-1]
		s.s = s.s[:len(s.s)-1]

		return top
	}

	return 0
}

func (s myStack) top() int {
	return s.s[len(s.s)-1]
}

func (s myStack) isEmpty() bool {
	return len(s.s) == 0
}

type MyQueue struct {
	s1    myStack
	s2    myStack
	front int
}

func Constructor() MyQueue {
	return MyQueue{
		s1:    constructor(),
		s2:    constructor(),
		front: 0,
	}
}

// time complexity: each element in stack has 2 push, 2 pop => 4n
// new element has 1 push, 1 assign => 2
// total is O(4n+2) => O(n)
func (this *MyQueue) Push(x int) {
	if this.s1.isEmpty() {
		this.front = x
	}

	for !this.s1.isEmpty() {
		top := this.s1.pop()
		this.s2.push(top)
	}

	this.s1.push(x)

	for !this.s2.isEmpty() {
		top := this.s2.pop()
		this.s1.push(top)
	}
}

// time complexity: O(1)
func (this *MyQueue) Pop() int {
	top := this.s1.pop()
	if !this.s1.isEmpty() {
		this.front = this.s1.top()
	}

	return top
}

func (this *MyQueue) Peek() int {
	return this.front
}

func (this *MyQueue) Empty() bool {
	return this.s1.isEmpty()
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
