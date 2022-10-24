package datastructure

type MyQueue struct {
	head   *node
	tail   *node
	length int
}

type node struct {
	val  interface{}
	next *node
}

func NewQue() MyQueue {
	return MyQueue{}
}

func (q MyQueue) Front() interface{} {
	if q.head != nil {
		return q.head.val
	}

	return nil
}

func (q MyQueue) IsEmpty() bool {
	return q.length == 0
}

func (q *MyQueue) Enqueue(v interface{}) {
	if q.head == nil {
		q.head = new(node)
		q.head.val = v
		q.tail = q.head
	} else {
		newNode := new(node)
		newNode.val = v
		q.tail.next = newNode
		q.tail = q.tail.next
	}

	q.length++
}

func (q *MyQueue) Dequeue() interface{} {
	if q.length > 0 {
		tmp := q.head
		q.head = q.head.next
		q.length--
		tmp.next = nil
		return tmp.val
	}

	return nil
}

func (q MyQueue) Length() int {
	return q.length
}
