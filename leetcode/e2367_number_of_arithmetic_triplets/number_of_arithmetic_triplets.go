package numberofarithmetictriplets

// two pointers
// Time: O(n^2)
func arithmeticTriplets(nums []int, diff int) int {
	if len(nums) < 3 {
		return 0
	}

	totalCount := 0
	for i := 0; i < len(nums); i++ {
		slow, fast := i, i+1
		count := 0
		for fast < len(nums) {
			if nums[fast]-nums[slow] == diff {
				count++
				slow = fast

				if count == 2 {
					totalCount++
					break
				}
			}

			fast++
		}
	}

	return totalCount
}

// hash
// Time: O(n), Space: O(n)
func arithmeticTripletsHash(nums []int, diff int) int {
	if len(nums) < 3 {
		return 0
	}

	totalCount := 0
	hash := make(map[int]bool)
	for _, v := range nums {
		if v >= diff*2 {
			if (hash[v-diff] == true) && (hash[v-2*diff] == true) {
				totalCount++
			}
		}

		hash[v] = true
	}

	return totalCount
}
