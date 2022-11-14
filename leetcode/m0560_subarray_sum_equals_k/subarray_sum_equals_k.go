package subarraysumequalsk

// Time complexity: O(n)
// Space complexity: O(n)
func subarraySum(nums []int, k int) int {
	hash := make(map[int]int)
	hash[0] = 1
	sum, ans := 0, 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if k := hash[sum-k]; k > 0 {
			ans += k
		}

		hash[sum]++
	}

	return ans
}
