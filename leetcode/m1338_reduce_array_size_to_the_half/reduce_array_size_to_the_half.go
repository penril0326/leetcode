package reducearraysizetothehalf

import "container/heap"

// Time: O(n*logn)
// Space: O(k), k is how many different number in the arr. The worst case is O(n).
func minSetSize(arr []int) int {
	h := map[int]int{}
	for _, v := range arr {
		h[v]++
	}

	q := &pq{}
	for _, count := range h {
		*q = append(*q, count)
	}

	heap.Init(q)
	remove := 0
	removeCount := 0
	for q.Len() > 0 {
		count := heap.Pop(q).(int)
		remove += count
		removeCount++
		if remove >= (len(arr) / 2) {
			break
		}
	}

	return removeCount
}

type pq []int

func (q pq) Len() int {
	return len(q)
}

func (q pq) Less(i, j int) bool {
	return q[i] > q[j]
}

func (q pq) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *pq) Push(v interface{}) {
	*q = append(*q, v.(int))
}

func (q *pq) Pop() interface{} {
	if q.Len() == 0 {
		return nil
	}

	lastIdx := q.Len() - 1
	root := (*q)[lastIdx]
	*q = (*q)[:lastIdx]

	return root
}
