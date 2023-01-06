package reverseonlyletters

import (
	"bytes"
	"practice/data_structure/stack"
)

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

// Using stack
// Time: O(N)
// Space: O(N)
func reverseOnlyLettersStack(s string) string {
	st := stack.NewStack()
	for i := 0; i < len(s); i++ {
		if isLetter(s[i]) {
			st.Push(s[i])
		}
	}

	var ans bytes.Buffer
	for i := 0; i < len(s); i++ {
		if isLetter(s[i]) {
			ans.WriteByte(st.Pop().(byte))
		} else {
			ans.WriteByte(s[i])
		}
	}

	return ans.String()
}

func isLetter(c byte) bool {
	if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
		return true
	}

	return false
}
