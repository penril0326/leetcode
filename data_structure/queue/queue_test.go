package queue

import (
	"log"
	"reflect"
	"testing"
)

func TestNewQue(t *testing.T) {
	de := NewQue()
	de.Enqueue(1)
	assert(de.Front(), 1)
	assert(de.IsEmpty(), false)
	assert(de.Length(), 1)

	de.Enqueue(2)
	assert(de.Front(), 1)
	assert(de.IsEmpty(), false)
	assert(de.Length(), 2)

	de.Enqueue(3)
	assert(de.Front(), 1)
	assert(de.IsEmpty(), false)
	assert(de.Length(), 3)

	assert(de.Front(), 1)
	assert(de.Dequeue(), 1)
	assert(de.IsEmpty(), false)
	assert(de.Length(), 2)

	assert(de.Front(), 2)
	assert(de.Dequeue(), 2)
	assert(de.IsEmpty(), false)
	assert(de.Length(), 1)

	assert(de.Length(), 1)
	assert(de.Dequeue(), 3)
	assert(de.Front(), nil)
	assert(de.IsEmpty(), true)
	assert(de.Length(), 0)

	assert(de.Dequeue(), nil)
	assert(de.Front(), nil)
	assert(de.IsEmpty(), true)
	assert(de.Length(), 0)
}

func assert(src, target interface{}) bool {
	if !reflect.DeepEqual(src, target) {
		log.Fatalf("src = %v, target = %v\n", src, target)
	}

	return true
}
