package longestcommonprefix

// without use any strings lib
// time complexity is O(S), which S is sum of characters in the strs
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}

	shortestStr := strs[0]
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < len(shortestStr) {
			shortestStr = strs[i]
		}
	}

	finalPosition := -1
	for i := 0; i < len(shortestStr); i++ {
		sameFlag := true
		for _, str := range strs {
			if str == shortestStr {
				continue
			}

			if str[i] != shortestStr[i] {
				sameFlag = false
				break
			}
		}

		if sameFlag == false {
			break
		} else {
			finalPosition++
		}
	}

	if finalPosition == -1 {
		return ""
	}

	return shortestStr[:finalPosition+1]
}
