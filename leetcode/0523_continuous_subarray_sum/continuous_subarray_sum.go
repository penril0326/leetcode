package continuous_subarray_sum

func checkSubarraySum(nums []int, k int) bool {
	length := len(nums)
	if 2 > length {
		return false
	}

	dp := make(map[int]int)
	dp[0] = -1

	sum := 0
	for i := 0; i < length; i++ {
		sum += nums[i]
		if 0 != k {
			sum = sum % k
		}

		if value, isExist := dp[sum]; true == isExist {
			if 1 < (i - value) {
				return true
			}
		} else {
			dp[sum] = i
		}
	}

	return false
}
