package reorganize_string

import (
	"sort"
)

type charCount struct {
	r     string
	count int
}

func reorganizeString(s string) string {
	// 初始化結構，先把a～z放好
	charCountList := make([]charCount, 26)
	for i := 0; i < 26; i++ {
		charCountList[i] = charCount{
			r:     string(i + 'a'),
			count: 0,
		}
	}

	// 計算a～z每個字母出現次數
	for _, r := range s {
		charCountList[r-'a'].count++

		// 若某個字母出現次數超過長度的一半，一定會有兩個一樣的字母相鄰
		// 等於長度一半的話可以
		if charCountList[r-'a'].count > (len(s)+1)/2 {
			return ""
		}
	}

	// 排序，次數多到少
	sort.Slice(charCountList, func(i, j int) bool {
		return charCountList[i].count > charCountList[j].count
	})

	result := make([]byte, len(s))
	idx := 0
	for i := 0; i < 26 && charCountList[i].count > 0; i++ {
		char := charCountList[i].r

		for charCountList[i].count > 0 {
			// 第一個字母先塞進result的第一個位置
			result[idx] = []byte(char)[0]
			charCountList[i].count--

			// 每次跳一格，若到底則從1開始
			// ex: "aaabbcc" -> a_a_a_b
			//                   b c c
			idx += 2
			if idx >= len(s) {
				idx = 1
			}
		}
	}

	return string(result)
}
