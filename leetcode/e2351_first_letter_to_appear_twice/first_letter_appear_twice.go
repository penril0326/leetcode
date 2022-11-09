package firstlettertoappeartwice

// Time complexity: O(1)
// Space complexity: O(N)
func repeatedCharacter(s string) byte {
	hash := [26]bool{}
	for _, r := range s {
		if hash[r-'a'] {
			return byte(r)
		}

		hash[r-'a'] = true
	}

	return 0
}
