package datastructure

type myQueue struct {
	head   *node
	tail   *node
	length int
}

type node struct {
	val  interface{}
	next *node
}

func NewQue() myQueue {
	return myQueue{}
}

func (q myQueue) Front() interface{} {
	if q.head != nil {
		return q.head.val
	}

	return nil
}

func (q myQueue) IsEmpty() bool {
	return q.length == 0
}

func (q *myQueue) Enqueue(v interface{}) {
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

func (q *myQueue) Dequeue() interface{} {
	if q.length > 0 {
		tmp := q.head
		q.head = q.head.next
		q.length--
		return tmp.val
	}

	return nil
}

func (q myQueue) Length() int {
	return q.length
}
