package firstuniquecharacterinastring

func firstUniqChar(s string) int {
	countingTable := make(map[rune]int)
	for _, r := range s {
		countingTable[r]++
	}

	for idx, r := range s {
		if countingTable[r] == 1 {
			return idx
		}
	}

	return -1
}
