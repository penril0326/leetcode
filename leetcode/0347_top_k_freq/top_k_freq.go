package top_k_freq

import (
	"container/heap"
)

type priorityQ []item

type item struct {
	value int
	count int
}

func topKFrequent(nums []int, k int) []int {
	countMap := map[int]int{}
	for _, v := range nums {
		countMap[v]++
	}

	temp := priorityQ{}
	for key, value := range countMap {
		heap.Push(&temp, item{
			value: key,
			count: value,
		})
	}

	result := []int{}
	for len(result) < k {
		i := heap.Pop(&temp).(item)
		result = append(result, i.value)
	}

	return result
}

func (q priorityQ) Len() int {
	return len(q)
}

func (q priorityQ) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q priorityQ) Less(i, j int) bool {
	return q[i].count > q[j].count
}

func (q *priorityQ) Push(v interface{}) {
	value := v.(item)
	*q = append(*q, value)
}

func (q *priorityQ) Pop() interface{} {
	length := len(*q)
	item := (*q)[length-1]
	*q = (*q)[:length-1]

	return item
}
