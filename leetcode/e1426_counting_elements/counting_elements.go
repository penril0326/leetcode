package countingelements

// Time complexity: O(N)
// Space complexity: O(N)
func countElements(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	hash := make(map[int]bool)
	for _, n := range arr {
		hash[n] = true
	}

	count := 0
	for _, n := range arr {
		if hash[n+1] {
			count++
		}
	}

	return count
}
