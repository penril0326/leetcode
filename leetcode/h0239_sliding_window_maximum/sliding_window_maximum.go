package slidingwindowmaximum

import "practice/data_structure/queue"

// Time complexity: O(n)
// Space complexity: O(n)
func maxSlidingWindow(nums []int, k int) []int {
	de := queue.NewDeque()
	result := make([]int, len(nums)-k+1)

	for i := 0; i < len(nums); i++ {
		for !de.IsEmpty() && (nums[de.Back().(int)] < nums[i]) {
			de.PopBack()
		}

		de.PushBack(i)

		if (i - de.Front().(int)) == k {
			de.PopFront()
		}

		if i >= (k - 1) {
			result[i-k+1] = nums[de.Front().(int)]
		}
	}

	return result
}
