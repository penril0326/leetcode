package group_anagrams

import "sort"

type runeList []rune

// solution 1: use sort
func groupAnagrams(strs []string) [][]string {
	wordMap := map[string][]string{}
	for _, str := range strs {
		rs := []rune(str)
		sort.Sort(runeList(rs))

		strList := wordMap[string(rs)]
		strList = append(strList, str)
		wordMap[string(rs)] = strList
	}

	result := [][]string{}
	for _, strList := range wordMap {
		result = append(result, strList)
	}

	return result
}

func (r runeList) Len() int {
	return len(r)
}

func (r runeList) Less(i, j int) bool {
	return r[i] < r[j]
}

func (r runeList) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

// solution 2: without sort
func groupAnagrams2(strs []string) [][]string {
	wordMap := map[string][]string{}
	for _, str := range strs {
		digitCountList := make([]int, 26)
		for _, c := range str {
			digitCountList[c-'a']++
		}

		wordKey := ""
		for i := 0; i < 26; i++ {
			if 0 == digitCountList[i] {
				continue
			}

			wordKey = wordKey + string(rune(i+'a')) + string(rune(digitCountList[i]))
		}

		strList := wordMap[wordKey]
		strList = append(strList, str)
		wordMap[wordKey] = strList
	}

	result := [][]string{}
	for _, strList := range wordMap {
		result = append(result, strList)
	}

	return result
}
