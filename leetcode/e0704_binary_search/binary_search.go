package binarysearch

// Time: O(logn)
// Space: O(1)
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) >> 1
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return -1
}
