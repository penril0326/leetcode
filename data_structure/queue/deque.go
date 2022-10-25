package queue

import "fmt"

type MyDeque interface {
	IsEmpty() bool
	Front() interface{}
	Back() interface{}
	PushFront(v interface{})
	PushBack(v interface{}) // Enqueue
	PopFront() interface{}  // Dequeue
	PopBack() interface{}
	Length() int
}

type deque struct {
	head   *node
	tail   *node
	length int
}

func NewDeque() MyDeque {
	return &deque{
		head: nil,
		tail: nil,
	}
}

func (de deque) IsEmpty() bool {
	return de.length == 0
}

func (de deque) Front() interface{} {
	if !de.IsEmpty() {
		return de.head.val
	}

	return nil
}

func (de deque) Back() interface{} {
	if !de.IsEmpty() {
		return de.tail.val
	}

	return nil
}

func (de *deque) PushFront(v interface{}) {
	newNode := new(node)
	newNode.val = v

	if de.head == nil {
		de.tail = newNode
	} else {
		newNode.next = de.head
		de.head.prev = newNode
	}

	de.head = newNode
	de.length++
}

func (de *deque) PushBack(v interface{}) {
	newNode := new(node)
	newNode.val = v

	if de.head == nil {
		de.head = newNode
	} else {
		newNode.prev = de.tail
		de.tail.next = newNode
	}

	de.tail = newNode
	de.length++
}

func (de *deque) PopFront() interface{} {
	if de.head != nil {
		head := de.head
		de.head = de.head.next
		head.next = nil

		if de.head == nil {
			de.tail = nil
		} else {
			de.head.prev = nil
		}

		de.length--

		return head.val
	}

	return nil
}

func (de *deque) PopBack() interface{} {
	if de.tail != nil {
		tail := de.tail
		de.tail = de.tail.prev
		tail.prev = nil

		if de.tail == nil {
			de.head = nil
		} else {
			de.tail.next = nil
		}

		de.length--

		return tail.val
	}

	return nil
}

func (de deque) Length() int {
	return de.length
}

func (de deque) Print() {
	p := de.head
	for p.next != nil {
		fmt.Printf("%d ", p.val)
		p = p.next
	}
	fmt.Println()
}
