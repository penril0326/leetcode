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

// Time complexity: O(N)
// Space complexity: O(1)
func longestOnesOptimized(nums []int, k int) int {
	zero, left := 0, 0
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zero++
		}

		// first occur zero > k which mean now the window size is the maximum
		// now we move left every time, if we found zero and throw out, we fix left again.
		// If right still < len(nums), the window size only grow up.
		if zero > k {
			if nums[left] == 0 {
				zero--
			}

			left++
		}
	}

	return len(nums) - left
}
