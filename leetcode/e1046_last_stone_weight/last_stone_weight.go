package laststoneweight

import "container/heap"

// Time: O(nLogn), n is total of stones
// Space: O(n)
func lastStoneWeight(stones []int) int {
	h := &weights{}
	for _, stone := range stones {
		*h = append(*h, stone)
	}

	heap.Init(h)

	for h.Len() > 1 {
		x, y := heap.Pop(h).(int), heap.Pop(h).(int)
		if x != y {
			heap.Push(h, x-y)
		}
	}

	if h.Len() > 0 {
		return heap.Pop(h).(int)
	}

	return 0
}

type weights []int

func (w weights) Len() int {
	return len(w)
}

func (w weights) Less(i, j int) bool {
	return w[i] >= w[j]
}

func (w weights) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

func (w *weights) Push(x interface{}) {
	*w = append(*w, x.(int))
}

func (w *weights) Pop() interface{} {
	len := w.Len()
	v := (*w)[len-1]
	*w = (*w)[:len-1]
	return v
}
