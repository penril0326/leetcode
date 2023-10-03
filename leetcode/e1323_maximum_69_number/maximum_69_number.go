package maximum69number

import (
	"strconv"
	"strings"
)

// Time: O(L), L is number of digit in num. (in this case is 4)
// Space: O(L)
func maximum69Number(num int) int {
	s := strconv.Itoa(num)
	var newS strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] == '6' {
			newS.WriteString(s[:i])
			newS.WriteByte('9')
			newS.WriteString(s[i+1:])

			newInt, _ := strconv.Atoi(newS.String())
			return newInt
		}
	}

	return num
}
