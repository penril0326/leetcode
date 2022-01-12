package lengthoflastword

func lengthOfLastWord(s string) int {
	length := len(s)
	if length == 0 {
		return 0
	}

	count := 0
	for i := length - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if count > 0 {
				break
			}
		} else {
			count++
		}
	}

	return count
}
