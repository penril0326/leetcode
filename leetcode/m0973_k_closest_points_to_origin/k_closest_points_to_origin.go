package kclosestpointstoorigin

import "container/heap"

// Time: O((n+k)*logk)
// Space: O(k)
func kClosest(points [][]int, k int) [][]int {
	q := &pq{}
	for _, point := range points {
		heap.Push(q, point)

		if q.Len() > k {
			heap.Pop(q)
		}
	}

	ans := [][]int{}
	for q.Len() > 0 {
		ans = append(ans, heap.Pop(q).([]int))
	}

	return ans
}

type pq [][]int

func (q pq) Len() int {
	return len(q)
}

func (q pq) Less(i, j int) bool {
	xi, yi := q[i][0], q[i][1]
	xj, yj := q[j][0], q[j][1]

	return (xi*xi + yi*yi) > (xj*xj + yj*yj)
}

func (q pq) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *pq) Push(v interface{}) {
	*q = append(*q, v.([]int))
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
