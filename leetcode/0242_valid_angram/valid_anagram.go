package valid_anagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count1 := map[rune]int{}
	count2 := map[rune]int{}

	for _, r := range s {
		count1[r]++
	}

	for _, r := range t {
		count2[r]++
	}

	for r, count1 := range count1 {
		if count2, isExist := count2[r]; (isExist == false) || (count1 != count2) {
			return false
		}
	}

	return true
}
