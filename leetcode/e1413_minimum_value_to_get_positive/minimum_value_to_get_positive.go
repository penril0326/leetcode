package minimumvaluetogetpositive

// Time complexity: O(N)
// Space complexity: O(1)
func minStartValue(nums []int) int {
	prefixSum, minimum := 0, 0
	for i := 0; i < len(nums); i++ {
		prefixSum += nums[i]

		if minimum > prefixSum {
			minimum = prefixSum
		}
	}

	return 1 - minimum
}
