package checkifthesentenceispangram

// Time complexity: O(N)
// Space complexity: O(1)
func checkIfPangram(sentence string) bool {
	alphabet := make(map[rune]bool)
	for _, r := range sentence {
		alphabet[r] = true
	}

	return len(alphabet) == 26
}
