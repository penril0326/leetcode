package minimumcostconnectsticks

import "container/heap"

// Time: O(NlogN), N is size of sticks
// Space: O(N)
func connectSticks(sticks []int) int {
	if len(sticks) <= 1 {
		return 0
	}

	mh := &minHeap{}
	*mh = append(*mh, sticks...)

	heap.Init(mh)
	totalCost := 0
	for mh.Len() > 1 {
		first, second := heap.Pop(mh).(int), heap.Pop(mh).(int)
		cost := first + second
		totalCost += cost
		heap.Push(mh, cost)
	}

	return totalCost
}

type minHeap []int

func (m minHeap) Len() int {
	return len(m)
}

func (m minHeap) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m minHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *minHeap) Push(v interface{}) {
	*m = append(*m, v.(int))
}

func (m *minHeap) Pop() interface{} {
	if m.Len() == 0 {
		return nil
	}

	lastIdx := m.Len() - 1
	root := (*m)[lastIdx]
	*m = (*m)[:lastIdx]

	return root
}
