package searchinsertposition

// Time: O(logn)
// Space: O(1)
func searchInsert(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	// if target bigger than last element just insert to the end
	if target > nums[length-1] {
		return length
	}

	// binary search
	low, high := 0, length-1
	for low < high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid
		} else {
			low = mid + 1
		}
	}

	// finally low is equal high so is ok to return one of them
	return high
}
