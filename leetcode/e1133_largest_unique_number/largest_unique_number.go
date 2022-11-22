package largestuniquenumber

// Time complexity: O(n)
// Space complexity: O(n)
func largestUniqueNumber(nums []int) int {
	hash := make(map[int]int)
	max := 0
	for _, n := range nums {
		if n > max {
			max = n
		}

		hash[n]++
	}

	for i := max; i >= 0; i-- {
		if hash[i] == 1 {
			return i
		}
	}

	return -1
}
