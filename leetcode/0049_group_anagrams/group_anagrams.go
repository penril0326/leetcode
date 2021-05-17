package group_anagrams

import (
	"sort"
)

type runeList []rune

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
