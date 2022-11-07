package maxconsecutiveonesiii

// Time complexity: O(N)
// Space complexity: O(1)
func longestOnes(nums []int, k int) int {
	zero, left, max := 0, 0, 0
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zero++
		}

		for zero > k {
			if nums[left] == 0 {
				zero--
			}

			left++
		}

		curr := right - left + 1
		if curr > max {
			max = curr
		}
	}

	return max
}
