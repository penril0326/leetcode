package searchinsertposition

// Time: O(logn)
// Space: O(1)
func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	// in this case, the relation of left ad right will be: right < left
	// at this point, nums[right] < target < nums[left]
	// so we just return left pointer

	return l
}
