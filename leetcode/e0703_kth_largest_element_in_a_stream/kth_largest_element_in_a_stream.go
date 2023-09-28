package kthlargestelementinastream

import (
	"container/heap"
)

type KthLargest struct {
	vals *myHeap
	k    int
}

// O(N*logk), N is size of nums
func Constructor(k int, nums []int) KthLargest {
	kth := KthLargest{
		vals: &myHeap{},
		k:    k,
	}

	heap.Init(kth.vals)
	for _, v := range nums {
		kth.Add(v)
	}

	return kth
}

// O(M*logk), which M is the call times for add()
func (this *KthLargest) Add(val int) int {
	// O(logk)
	heap.Push(this.vals, val)

	if this.vals.Len() > this.k {
		// O(logk)
		heap.Pop(this.vals)
	}

	return this.vals.Top()
}

type myHeap []int

func (h myHeap) Len() int {
	return len(h)
}

func (h myHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h myHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *myHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func (h *myHeap) Pop() interface{} {
	if h.Len() == 0 {
		panic("Pop no element in heap")
	}

	lastIdx := h.Len() - 1
	root := (*h)[lastIdx]
	*h = (*h)[:lastIdx]

	return root
}

func (h myHeap) Top() int {
	if h.Len() == 0 {
		panic("Top no element in heap")
	}

	return h[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
