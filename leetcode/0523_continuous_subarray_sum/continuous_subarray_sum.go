package continuous_subarray_sum

func checkSubarraySum(nums []int, k int) bool {
	length := len(nums)
	if 2 > length {
		return false
	}

	// key = remainder, value = index of nums
	dp := make(map[int]int)
	dp[0] = -1

	sum := 0
	for i := 0; i < length; i++ {
		sum += nums[i]
		t := 0
		if 0 != k {
			t = sum % k
		}

		if index, isExist := dp[t]; true == isExist {
			if 1 < (i - index) {
				return true
			}
		} else {
			dp[t] = i
		}
	}

	return false
}
