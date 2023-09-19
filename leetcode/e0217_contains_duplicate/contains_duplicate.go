package containsduplicate

// Time: O(n)
// Space: O(n)
func containsDuplicate(nums []int) bool {
	hash := make(map[int]int)
	for _, value := range nums {
		if _, isExist := hash[value]; isExist {
			return true
		}

		hash[value]++
	}

	return false
}
