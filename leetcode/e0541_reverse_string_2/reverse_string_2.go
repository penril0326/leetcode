package reversestring2

// Time: O(n)
// Space: O(n)
func reverseStr(s string, k int) string {
	b := []byte(s)
	for cur := 0; cur < len(s); cur += 2 * k {
		i, j := cur, min(cur+k-1, len(s)-1)
		for i < j {
			b[i], b[j] = b[j], b[i]
			i++
			j--
		}
	}

	return string(b)
}

func min(x, y int) int {
	if x > y {
		return y
	}

	return x
}
