package longest_word

import "sort"

type stringList []string

// use sort
func findLongestWord(s string, dictionary []string) string {
	d := stringList{}
	d = append(d, dictionary...)
	sort.Sort(d)

	resultIdx := -1
	maxLength := 0
	for idx, word := range d {
		if len(word) > len(s) {
			continue
		}

		wordIdx := 0
		sIdx := 0
		for (wordIdx < len(word)) && (sIdx < len(s)) {
			if word[wordIdx] == s[sIdx] {
				wordIdx++
			}

			sIdx++
		}

		if (wordIdx == len(word)) && (len(d[idx]) > maxLength) {
			resultIdx = idx
			maxLength = len(d[idx])
		}
	}

	if resultIdx == -1 {
		return ""
	}

	return d[resultIdx]
}

// without sort
func findLongestWord2(s string, dictionary []string) string {
	result := ""
	for _, word := range dictionary {
		if len(word) > len(s) {
			continue
		}

		wordIdx := 0
		for _, r := range s {
			if rune(word[wordIdx]) == r {
				wordIdx++

				if wordIdx == len(word) {
					break
				}
			}
		}

		if (wordIdx == len(word)) && (len(word) >= len(result)) {
			if len(word) > len(result) {
				result = word
			} else {
				result = smallString(word, result)
			}
		}
	}

	return result
}

func smallString(s1, s2 string) string {
	if s1 < s2 {
		return s1
	}

	return s2
}

func (sl stringList) Len() int {
	return len(sl)
}

func (sl stringList) Less(i, j int) bool {
	shortLength := len(sl[i])
	if len(sl[j]) < shortLength {
		shortLength = len(sl[j])
	}

	for idx := 0; idx < shortLength; idx++ {
		if sl[i][idx] < sl[j][idx] {
			return true
		} else if sl[i][idx] > sl[j][idx] {
			return false
		}
	}

	return len(sl[i]) < len(sl[j])
}

func (sl stringList) Swap(i, j int) {
	sl[i], sl[j] = sl[j], sl[i]
}
