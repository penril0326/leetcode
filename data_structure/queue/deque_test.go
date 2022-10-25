package queue

import (
	"testing"
)

func TestNewDeque(t *testing.T) {
	de := NewDeque()
	de.PushBack(1)
	assert(de.Front(), 1)
	assert(de.Back(), 1)
	assert(de.IsEmpty(), false)
	assert(de.Length(), 1)

	de.PushBack(2)
	assert(de.Front(), 1)
	assert(de.Back(), 2)
	assert(de.IsEmpty(), false)
	assert(de.Length(), 2)

	de.PushBack(3)
	assert(de.Front(), 1)
	assert(de.Back(), 3)
	assert(de.IsEmpty(), false)
	assert(de.Length(), 3)

	de.PopFront()
	assert(de.Front(), 2)
	assert(de.Back(), 3)
	assert(de.IsEmpty(), false)
	assert(de.Length(), 2)

	de.PopBack()
	assert(de.Front(), 2)
	assert(de.Back(), 2)
	assert(de.IsEmpty(), false)
	assert(de.Length(), 1)

	de.PopFront()
	assert(de.Front(), nil)
	assert(de.Back(), nil)
	assert(de.IsEmpty(), true)
	assert(de.Length(), 0)

	de.PushFront(5)
	assert(de.Front(), 5)
	assert(de.Back(), 5)
	assert(de.IsEmpty(), false)
	assert(de.Length(), 1)

	de.PushBack(6)
	assert(de.Front(), 5)
	assert(de.Back(), 6)
	assert(de.IsEmpty(), false)
	assert(de.Length(), 2)

	de.PopFront()
	de.PopBack()
	assert(de.Front(), nil)
	assert(de.Back(), nil)
	assert(de.IsEmpty(), true)
	assert(de.Length(), 0)

	de.PopFront()
	de.PopBack()
	assert(de.Front(), nil)
	assert(de.Back(), nil)
	assert(de.IsEmpty(), true)
	assert(de.Length(), 0)
}
