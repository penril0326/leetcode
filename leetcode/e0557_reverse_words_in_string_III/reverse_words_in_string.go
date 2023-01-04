package reversewordsinstringiii

import (
	"bytes"
	"strings"
)

// Time: O(n), every single char is traversed twice
// Space: O(1)
func reverseWords(s string) string {
	var ans strings.Builder
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			// use string builder and byte buffer to concate string is O(1)
			// use '+' is cost O(N)
			tmp := reverseString(s[start:i])
			ans.WriteString(tmp)
			ans.WriteString(" ")
			start = i + 1
		}
	}

	tmp := reverseString(s[start:])
	ans.WriteString(tmp)

	return ans.String()
}

func reverseString(s string) string {
	var ans bytes.Buffer
	for i := len(s) - 1; i >= 0; i-- {
		ans.WriteByte(s[i])
	}

	return ans.String()
}
