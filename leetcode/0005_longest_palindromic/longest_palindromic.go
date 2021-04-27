package longest_palindromic

func longestPalindrome(s string) string {
	resultEven := getLongestEvenPalindrome(s)
	resultOdd := getLongestOddPalindrome(s)

	if len(resultOdd) > len(resultEven) {
		return resultOdd
	}

	return resultEven
}

func getLongestEvenPalindrome(s string) string {
	maxLength := 0
	minLeft, maxRight := 0, 0
	for i := 0; i < len(s)-1; i++ {
		if len(s) > (i + 1) {
			if s[i] != s[i+1] {
				continue
			}
		}

		left, right := i-1, i+2
		isFind := false
		for (0 < left) && (len(s) > right) {
			if s[left] != s[right] {
				break
			}

			left--
			right++
			isFind = true
		}

		if (true == isFind) && (maxLength < (right - left)) {
			if 0 > left {
				left = 0
			}

			if len(s) <= right {
				right = len(s) - 1
			}

			maxLength = right - left + 1
			minLeft = left
			maxRight = right
		}
	}

	return s[minLeft:maxRight]
}

func getLongestOddPalindrome(s string) string {
	maxLength := 0
	minLeft, maxRight := 0, 0
	for i := 0; i < len(s); i++ {
		left, right := i, i
		isFind := false
		for (0 <= left) && (len(s) > right) && (s[left] == s[right]) {
			isFind = true
			left--
			right++
		}

		if (true == isFind) && (maxLength < (right - left)) {
			if 0 > left {
				left = 0
			}

			if len(s) <= right {
				right = len(s) - 1
			}

			maxLength = right - left + 1
			minLeft = left
			maxRight = right
		}
	}

	return s[minLeft:maxRight]
}
