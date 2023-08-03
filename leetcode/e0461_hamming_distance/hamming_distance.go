package hammingdistance

// Time: O(1)
// Space: O(1)
func hammingDistance(x int, y int) int {
	xor := x ^ y

	countOne := 0
	for xor != 0 {
		if xor%2 == 1 {
			countOne++
		}

		xor = xor >> 1
	}

	return countOne
}
