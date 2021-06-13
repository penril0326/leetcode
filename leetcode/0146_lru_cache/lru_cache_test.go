package lru_cache

import (
	"testing"
)

func TestLRU(t *testing.T) {
	// testCase1(t)
	// testCase2(t)
	testCase3(t)
}

func testCase1(t *testing.T) {
	lru := Constructor(2)

	lru.Put(1, 10)
	t.Logf("head = %v", lru.head.value)
	t.Logf("tail = %v", lru.tail.value)
	t.Log("-----")

	lru.Put(2, 20)
	t.Logf("head = %v", lru.head.value)
	t.Logf("tail = %v", lru.tail.value)
	t.Log("-----")

	t.Logf("%d", lru.Get(1))
	t.Logf("head = %v", lru.head.value)
	t.Logf("tail = %v", lru.tail.value)
	t.Log("-----")

	lru.Put(3, 30)
	t.Logf("head = %v", lru.head.value)
	t.Logf("tail = %v", lru.tail.value)
	t.Log("-----")

	t.Logf("%d", lru.Get(2))
	t.Logf("head = %v", lru.head.value)
	t.Logf("tail = %v", lru.tail.value)
	t.Log("-----")

	lru.Put(4, 40)
	t.Logf("head = %v", lru.head.value)
	t.Logf("tail = %v", lru.tail.value)
	t.Log("-----")

	t.Logf("%d", lru.Get(1))
	t.Logf("head = %v", lru.head.value)
	t.Logf("tail = %v", lru.tail.value)
	t.Log("-----")

	t.Logf("%d", lru.Get(3))
	t.Logf("head = %v", lru.head.value)
	t.Logf("tail = %v", lru.tail.value)
	t.Log("-----")

	t.Logf("%d", lru.Get(4))
	t.Logf("head = %v", lru.head.value)
	t.Logf("tail = %v", lru.tail.value)
}

func testCase2(t *testing.T) {
	lru := Constructor(1)

	lru.Put(2, 1)
	t.Logf("head = %v", lru.head.value)
	t.Logf("tail = %v", lru.tail.value)
	t.Log("-----")

	t.Logf("%d", lru.Get(2))
	t.Logf("head = %v", lru.head.value)
	t.Logf("tail = %v", lru.tail.value)
	t.Log("-----")

	lru.Put(3, 2)
	t.Logf("head = %v", lru.head.value)
	t.Logf("tail = %v", lru.tail.value)
	t.Log("-----")

	t.Logf("%d", lru.Get(2))
	t.Logf("head = %v", lru.head.value)
	t.Logf("tail = %v", lru.tail.value)
	t.Log("-----")

	t.Logf("%d", lru.Get(3))
	t.Logf("head = %v", lru.head.value)
	t.Logf("tail = %v", lru.tail.value)
	t.Log("-----")
}

func testCase3(t *testing.T) {
	lru := Constructor(10)

	lru.Put(10, 13)
	lru.Put(3, 17)
	lru.Put(6, 11)
	lru.Put(10, 5)
	lru.Put(9, 10)
	lru.Get(13)
	lru.Put(2, 19)
	lru.Get(2)
	lru.Get(3)
	lru.Put(5, 25)
	lru.Get(8)
	lru.Put(9, 22)
	lru.Put(5, 5)
	lru.Put(1, 30)
	lru.Get(11)
	lru.Put(9, 12)
	lru.Get(7)
	lru.Get(5)
	lru.Get(8)
	lru.Get(9)
	lru.Put(4, 30)
	lru.Put(9, 3)
	lru.Get(9)
	lru.Get(10)
	lru.Get(10)
	lru.Put(6, 14)
	lru.Put(3, 1)
	lru.Get(3)
	lru.Put(10, 11)
	lru.Get(8)
}
