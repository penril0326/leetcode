package numberofwaystosplitarray

// Time complexity: O(N)
// Space complexity: O(N)
func waysToSplitArray(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}

	prefixSum := make([]int, len(nums))
	prefixSum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i]
	}

	count := 0
	for i := 0; i < len(nums)-1; i++ {
		if prefixSum[i] >= (prefixSum[len(nums)-1] - prefixSum[i]) {
			count++
		}
	}

	return count
}
