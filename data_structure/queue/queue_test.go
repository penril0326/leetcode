package queue

import (
	"log"
	"reflect"
	"testing"
)

func TestNewQue(t *testing.T) {
	de := NewDeque()
	de.Enqueue(1)
	assert(de.Front(), 1)
	assert(de.Back(), 1)
	assert(de.IsEmpty(), false)
	assert(de.Length(), 1)

	de.Enqueue(2)
	assert(de.Front(), 1)
	assert(de.Back(), 2)
	assert(de.IsEmpty(), false)
	assert(de.Length(), 2)

	de.Enqueue(3)
	assert(de.Front(), 1)
	assert(de.Back(), 3)
	assert(de.IsEmpty(), false)
	assert(de.Length(), 3)

	de.Dequeue()
	assert(de.Front(), 2)
	assert(de.Back(), 3)
	assert(de.IsEmpty(), false)
	assert(de.Length(), 2)

	de.PopBack()
	assert(de.Front(), 2)
	assert(de.Back(), 2)
	assert(de.IsEmpty(), false)
	assert(de.Length(), 1)

	de.Dequeue()
	assert(de.Front(), nil)
	assert(de.Back(), nil)
	assert(de.IsEmpty(), true)
	assert(de.Length(), 0)

	de.PushFront(5)
	assert(de.Front(), 5)
	assert(de.Back(), 5)
	assert(de.IsEmpty(), false)
	assert(de.Length(), 1)

	de.Enqueue(6)
	assert(de.Front(), 5)
	assert(de.Back(), 6)
	assert(de.IsEmpty(), false)
	assert(de.Length(), 2)

	de.Dequeue()
	de.PopBack()
	assert(de.Front(), nil)
	assert(de.Back(), nil)
	assert(de.IsEmpty(), true)
	assert(de.Length(), 0)

	de.Dequeue()
	de.PopBack()
	assert(de.Front(), nil)
	assert(de.Back(), nil)
	assert(de.IsEmpty(), true)
	assert(de.Length(), 0)
}

func assert(src, target interface{}) bool {
	if !reflect.DeepEqual(src, target) {
		log.Fatalf("src = %v, target = %v\n", src, target)
	}

	return true
}
