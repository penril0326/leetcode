package countnumberofnicesubarrays

// Time complexity: O(N)
// Space complexity: O(N)
func numberOfSubarrays(nums []int, k int) int {
	hash := make(map[int]int)
	oddCount := 0
	ans := 0
	hash[0]++
	for i := 0; i < len(nums); i++ {
		oddCount += nums[i] % 2
		ans += hash[oddCount-k]
		hash[oddCount]++
	}

	return ans
}
