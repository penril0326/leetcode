package firstlettertoappeartwice

// Time complexity: O(1)
// Space complexity: O(N)
func repeatedCharacter(s string) byte {
	hash := make(map[rune]bool)
	for _, r := range s {
		if _, exist := hash[r]; exist {
			return byte(r)
		}

		hash[r] = true
	}

	return 0
}
