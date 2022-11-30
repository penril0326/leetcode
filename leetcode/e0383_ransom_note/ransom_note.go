package ransomnote

// Time complexity: O(m), which m is the length of magazine
// Space complexity: O(1), since the hash map only store a-z, there are 26 alphabets.
func canConstruct(ransomNote string, magazine string) bool {
	charCount := make(map[rune]int) // you can also use [26]int instead of hash map
	for _, r := range magazine {
		charCount[r]++
	}

	for _, r := range ransomNote {
		charCount[r]--
		if charCount[r] < 0 {
			return false
		}
	}

	return true
}
