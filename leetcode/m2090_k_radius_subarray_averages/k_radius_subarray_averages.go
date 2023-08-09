package kradiussubarrayaverages

// Prefix Sum
// Time: O(N)
// Space: O(N)
func getAverages(nums []int, k int) []int {
	length := len(nums)
	result := make([]int, length)
	preSum := make([]int, length)
	preSum[0] = nums[0]
	for i := 1; i < length; i++ {
		preSum[i] = preSum[i-1] + nums[i]
	}

	curSum := 0
	for i := 0; i < length; i++ {
		curSum += nums[i]
		if (i-k) < 0 || (i+k) >= length {
			result[i] = -1
			continue
		}

		if (i - k) == 0 {
			result[i] = preSum[i+k] / (2*k + 1)
		} else {
			result[i] = (preSum[i+k] - preSum[i-k-1]) / (2*k + 1)
		}
	}

	return result
}
