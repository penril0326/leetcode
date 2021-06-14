package kth_largest

func findKthLargest(nums []int, k int) int {
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
