package findkclosestelements

import (
	"container/heap"
	"sort"
)

// Time: O(log(n-k))
// Space: O(1)
func findClosestElements(arr []int, k int, x int) []int {
	l, r := 0, len(arr)-k
	for l < r {
		mid := (l + r) >> 1
		if x-arr[mid] > arr[mid+k]-x {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return arr[l : l+k]
}

// Time: O((N+k)*logk), N is size of arr
// Space: O(k)
func findClosestElementsHeap(arr []int, k int, x int) []int {
	q := &pq{}
	for _, v := range arr {
		temp := [2]int{abs(v - x), v}
		heap.Push(q, temp)

		if q.Len() > k {
			heap.Pop(q)
		}
	}

	ans := make([]int, k)
	for i := 0; i < k; i++ {
		top := heap.Pop(q).([2]int)
		ans[i] = top[1]
	}

	sort.Ints(ans)

	return ans
}

func isClosest(a, b [2]int) bool {
	if (a[0] < b[0]) || ((a[0] == b[0]) && (a[1] < b[1])) {
		return true
	}

	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

// pq[i][0] is distance, pq[i][1] is value
type pq [][2]int

func (q pq) Len() int {
	return len(q)
}

func (q pq) Less(i, j int) bool {
	return !isClosest(q[i], q[j])
}

func (q pq) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *pq) Push(v interface{}) {
	*q = append(*q, v.([2]int))
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
