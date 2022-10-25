package queue

type MyQueue interface {
	IsEmpty() bool
	Front() interface{}
	Enqueue(v interface{})
	Dequeue() interface{}
	Length() int
}

type node struct {
	val  interface{}
	next *node
	prev *node
}

type myQueue struct {
	head   *node
	tail   *node
	length int
}

func NewQue() MyQueue {
	return &myQueue{
		head: nil,
		tail: nil,
	}
}

func (q myQueue) IsEmpty() bool {
	return q.length == 0
}

func (q myQueue) Front() interface{} {
	if q.head != nil {
		return q.head.val
	}

	return nil
}

func (q *myQueue) Enqueue(v interface{}) {
	newNode := new(node)
	newNode.val = v

	if q.head == nil {
		q.head = newNode
	} else {
		q.tail.next = newNode
	}

	q.tail = newNode

	q.length++
}

func (q *myQueue) Dequeue() interface{} {
	if q.head != nil {
		head := q.head
		q.head = q.head.next
		head.next = nil

		if q.head == nil {
			q.tail = nil
		}

		q.length--

		return head.val
	}

	return nil
}

func (q myQueue) Length() int {
	return q.length
}
