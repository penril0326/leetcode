package subarrayproductlessthank

// Time complexity: O(n)
// Space complexity: O(1)
func numSubarrayProductLessThanK(nums []int, k int) int {
	ans := 0
	product := 1
	left := 0
	for right := 0; right < len(nums); right++ {
		product *= nums[right]
		for (left <= right) && (product >= k) {
			product /= nums[left]
			left++
		}

		ans += (right - left + 1)
	}

	return ans
}
