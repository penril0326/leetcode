package maximumsubarray

func maxSubArray(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	} else if length == 1 {
		return nums[0]
	}

	// Kadane's algorithm
	// https://medium.com/@rsinghal757/kadanes-algorithm-dynamic-programming-how-and-why-does-it-work-3fd8849ed73d
	dp := make([]int, length)
	dp[0] = nums[0]
	result := dp[0]
	for i := 1; i < length; i++ {
		dp[i] = max(nums[i], nums[i]+dp[i-1])
		if dp[i] > result {
			result = dp[i]
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
