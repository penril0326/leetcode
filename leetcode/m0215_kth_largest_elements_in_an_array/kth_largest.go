package kth_largest

import "container/heap"

// Min-heap
// Time: O(n*logk)
// Space: O(k)
func findKthLargest(nums []int, k int) int {
	q := &pq{}
	for _, v := range nums {
		heap.Push(q, v)

		if q.Len() > k {
			heap.Pop(q)
		}
	}

	ans := heap.Pop(q).(int)
	return ans
}

type pq []int

func (q pq) Len() int {
	return len(q)
}

func (q pq) Less(i, j int) bool {
	return q[i] < q[j]
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

// Time: O(nlogn)
// Space: O(1), since we use in-place quick sort
func findKthLargestSort(nums []int, k int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		p := partition(nums, left, right)
		if p == k-1 {
			return nums[p]
		} else if p < k-1 {
			left = p + 1
		} else {
			right = p - 1
		}
	}

	return nums[k-1]
}

func partition(nums []int, left, right int) int {
	pivot := nums[right]
	i := left

	for j := left; j < right; j++ {
		if nums[j] > pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

	nums[i], nums[right] = nums[right], nums[i]
	return i
}
