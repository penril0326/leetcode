package numberofrecentcalls

import datastructure "practice/data_structure"

// Time complexity: O(1)
// Space complexity: O(1)
type RecentCounter struct {
	que datastructure.MyQueue
}

func Constructor() RecentCounter {
	return RecentCounter{
		que: datastructure.NewQue(),
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.que.Enqueue(t)

	min := t - 3000
	for !this.que.IsEmpty() && (this.que.Front().(int) < min) { // type assert will decrease performance
		this.que.Dequeue()
	}

	return this.que.Length()
}
