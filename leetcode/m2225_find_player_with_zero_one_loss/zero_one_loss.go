package findplayerwithzerooneloss

func findWinners(matches [][]int) [][]int {
	lose := make(map[int]int)
	win := make(map[int]int)
	maxUser := 0

	for _, match := range matches {
		maxUser = max(maxUser, match[0], match[1])
		lose[match[1]]++
		win[match[0]]++
	}

	noLose := make([]int, 0)
	oneLose := make([]int, 0)
	for player := 1; player <= maxUser; player++ {
		if loseCount, exist := lose[player]; !exist {
			if win[player] > 0 {
				noLose = append(noLose, player)
			}
		} else {
			if loseCount == 1 {
				oneLose = append(oneLose, player)
			}
		}
	}

	return [][]int{noLose, oneLose}
}

func max(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		} else {
			return c
		}
	} else {
		if b > c {
			return b
		} else {
			return c
		}
	}
}
