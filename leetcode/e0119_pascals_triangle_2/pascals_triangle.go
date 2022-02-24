package pascals_triangle2

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}

	if rowIndex == 1 {
		return []int{1, 1}
	}

	previous := []int{1, 1}
	for i := 2; i <= rowIndex; i++ {
		next := make([]int, i+1)
		next[0] = 1

		for j := 1; j < i; j++ {
			next[j] = previous[j-1] + previous[j]
		}

		next[i] = 1
		previous = next
	}

	return previous
}
