package counting_bits

// DP
func countingBits(n int) []int {
	dp := make([]int, n+1)

	for i := 0; i <= n; i++ {
		if i%2 == 0 {
			dp[i] = dp[i/2]
		} else {
			dp[i] = dp[i/2] + 1
		}
	}

	return dp
}

// general solution with O(nlogn)
func countingBitsG(n int) []int {
	result := make([]int, n+1)

	for i := 0; i <= n; i++ {
		countOne, temp := 0, i
		for temp != 0 {
			countOne += temp % 2
			temp /= 2
		}

		result[i] = countOne
	}

	return result
}

// special DP solution with O(n)
func countingBitsS(n int) []int {
	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		index := i & (i - 1)
		dp[i] = dp[index] + 1
	}

	return dp
}
