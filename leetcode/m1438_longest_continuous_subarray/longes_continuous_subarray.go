package longestcontinuoussubarray

import "practice/data_structure/queue"

func longestSubarray(nums []int, limit int) int {
	increasing := queue.NewDeque()
	decreasing := queue.NewDeque()
	leftIdx, finalAns := 0, 0

	for i := 0; i < len(nums); i++ {
		for !decreasing.IsEmpty() && (decreasing.Back().(int) < nums[i]) {
			decreasing.PopBack()
		}

		for !increasing.IsEmpty() && (increasing.Back().(int) > nums[i]) {
			increasing.PopBack()
		}

		decreasing.PushBack(nums[i])
		increasing.PushBack(nums[i])

		for (leftIdx < i) && (decreasing.Front().(int)-increasing.Front().(int) > limit) {
			if decreasing.Front().(int) == nums[leftIdx] {
				decreasing.PopFront()
			}

			if increasing.Front().(int) == nums[leftIdx] {
				increasing.PopFront()
			}

			leftIdx++
		}

		finalAns = max(finalAns, i-leftIdx+1)
	}

	return finalAns
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
