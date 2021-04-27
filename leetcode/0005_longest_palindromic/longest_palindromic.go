package longest_palindromic

func longestPalindrome(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		// odd case
		result = getPalindrome(s, result, i, i)

		// even case
		result = getPalindrome(s, result, i, i+1)
	}

	return result
}

func getPalindrome(s, res string, left, right int) string {
	sub := ""
	for (0 <= left) && (len(s) > right) && (s[left] == s[right]) {
		sub = s[left : right+1]
		left--
		right++
	}

	if len(res) < len(sub) {
		return sub
	}

	return res
}
