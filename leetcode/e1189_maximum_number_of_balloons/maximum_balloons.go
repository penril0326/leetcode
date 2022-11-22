package maximumnumberofballoons

// Time complexity: O(n)
// Space complexity: O(1)
func maxNumberOfBalloons(text string) int {
	if text == "" {
		return 0
	}

	hash := make(map[rune]int)
	for _, r := range text {
		switch r {
		case 'b', 'a', 'l', 'o', 'n':
			hash[r]++
		}
	}

	ans := hash['b']
	for key, value := range hash {
		count := value
		if (key == 'l') || (key == 'o') {
			count /= 2
		}

		if count < ans {
			ans = count
		}
	}

	return ans
}
