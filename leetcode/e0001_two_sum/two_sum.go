package twosum

// Time complexity: O(N)
// Space complexity: O(N)
func twoSum(nums []int, target int) []int {
	hash := map[int]int{}
	for idx, v := range nums {
		if i, exist := hash[target-v]; exist {
			return []int{i, idx}
		}

		hash[v] = idx
	}

	return nil
}
