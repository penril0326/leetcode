package numberofrecentcalls

import (
	"practice/data_structure/queue"
)

// Time complexity: O(1)
// Space complexity: O(1)
type RecentCounter struct {
	que queue.MyQueue
}

func Constructor() RecentCounter {
	return RecentCounter{
		que: queue.NewQue(),
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
