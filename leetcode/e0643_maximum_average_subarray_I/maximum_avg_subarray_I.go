package maximumaveragesubarrayi

// Time complexity: O(n)
// Space complexity: O(1)
func findMaxAverage(nums []int, k int) float64 {
	cur := 0
	for i := 0; i < k; i++ {
		cur += nums[i]
	}

	max := cur
	for right := k; right < len(nums); right++ {
		cur += nums[right] - nums[right-k]

		if cur > max {
			max = cur
		}
	}

	return float64(max) / float64(k)
}
