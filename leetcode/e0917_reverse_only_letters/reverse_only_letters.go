package reverseonlyletters

// Time: O(N)
// Space: O(N)
func reverseOnlyLetters(s string) string {
	ans := []byte(s)
	start, end := 0, len(s)-1
	for start < end {
		if !isLetter(ans[start]) {
			start++
		} else if !isLetter(ans[end]) {
			end--
		} else {
			ans[start], ans[end] = ans[end], ans[start]
			start++
			end--
		}
	}

	return string(ans)
}

func isLetter(c byte) bool {
	if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
		return true
	}

	return false
}
