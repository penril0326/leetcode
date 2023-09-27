package findmedianfromdatastream

import "container/heap"

type MedianFinder struct {
	min *minHeap
	max *maxHeap
}

func Constructor() MedianFinder {
	mf := MedianFinder{
		min: &minHeap{},
		max: &maxHeap{},
	}

	heap.Init(mf.min)
	heap.Init(mf.max)

	return mf
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.max, num)
	heap.Push(this.min, heap.Pop(this.max))
	if this.min.Len() > this.max.Len() {
		heap.Push(this.max, heap.Pop(this.min))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	median := 0.0
	maxRoot := (*this.max)[0]
	if this.max.Len() == this.min.Len() {
		minRoot := (*this.min)[0]
		median = float64(maxRoot+minRoot) / 2.0
	} else {
		median = float64(maxRoot)
	}

	return median
}

type minHeap []int

func (m minHeap) Len() int {
	return len(m)
}

func (m minHeap) Less(i, j int) bool {
	return m[i] <= m[j]
}

func (m minHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *minHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *minHeap) Pop() interface{} {
	n := m.Len()
	root := (*m)[n-1]
	*m = (*m)[:n-1]
	return root
}

type maxHeap []int

func (m maxHeap) Len() int {
	return len(m)
}

func (m maxHeap) Less(i, j int) bool {
	return m[i] >= m[j]
}

func (m maxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *maxHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *maxHeap) Pop() interface{} {
	n := m.Len()
	root := (*m)[n-1]
	*m = (*m)[:n-1]
	return root
}
