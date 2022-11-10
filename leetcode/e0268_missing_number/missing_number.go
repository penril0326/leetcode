package missingnumber

// Time complexity: O(N)
// Space complexity: O(N)
func missingNumber(nums []int) int {
	hash := make(map[int]bool)
	for _, n := range nums {
		hash[n] = true
	}

	for i := 0; i <= len(nums); i++ {
		if !hash[i] {
			return i
		}
	}

	return -1
}

// One pass, Gauss' Formula
// Time complexity: O(N)
// Space complexity: O(1)
func missingNumber2(nums []int) int {
	n := len(nums)
	total := n * (n + 1) / 2

	for i := 0; i < len(nums); i++ {
		total -= nums[i]
	}

	return total
}
