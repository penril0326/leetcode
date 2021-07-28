package count_binary_substrings

//00110011
func countBinarySubstrings(s string) int {
	if len(s) == 0 {
		return 0
	}

	countZero := 0
	countOne := 0
	result := 0

	if s[0] == '0' {
		countZero++
	} else if s[0] == '1' {
		countOne++
	}

	for i := 1; i < len(s); i++ {
		if s[i] == '0' {
			if s[i-1] != s[i] {
				countZero = 1
			} else {
				countZero++
			}

			if countOne >= countZero {
				result++
			}
		} else if s[i] == '1' {
			if s[i-1] != s[i] {
				countOne = 1
			} else {
				countOne++
			}

			if countZero >= countOne {
				result++
			}
		}
	}

	return result
}
