package howmanyapplescanyouputintothebucket

import "sort"

// Time: O(n*logn)
// Space: O(1)
func maxNumberOfApples(weight []int) int {
	sort.Ints(weight)
	total := 5000
	i := 0
	count := 0
	for (i < len(weight)) && (total > 0) {
		if total-weight[i] >= 0 {
			total -= weight[i]
			count++
		}

		i++
	}

	return count
}
