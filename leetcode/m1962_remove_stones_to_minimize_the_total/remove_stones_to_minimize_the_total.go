package removestonestominimizethetotal

import (
	"container/heap"
	"math"
)

// Time: O(k*logn+n), n is size of piles
// Space: O(n)
func minStoneSum(piles []int, k int) int {
	maxh := &maxHeap{}
	*maxh = append(*maxh, piles...)

	heap.Init(maxh)
	for i := k; i > 0; i-- {
		max := heap.Pop(maxh).(int)
		roundup := math.Ceil(float64(max) / 2.0)
		heap.Push(maxh, int(roundup))
	}

	sum := 0
	for maxh.Len() > 0 {
		sum += heap.Pop(maxh).(int)
	}

	return sum
}

type maxHeap []int

func (m maxHeap) Len() int {
	return len(m)
}

func (m maxHeap) Less(i, j int) bool {
	return m[i] > m[j]
}

func (m maxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *maxHeap) Push(v interface{}) {
	*m = append(*m, v.(int))
}

func (m *maxHeap) Pop() interface{} {
	if m.Len() == 0 {
		return nil
	}

	// old := *m
	lastIdx := m.Len() - 1
	x := (*m)[lastIdx]
	*m = (*m)[:lastIdx]

	return x
}
