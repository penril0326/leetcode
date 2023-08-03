package nimgame

// Time: O(1)
// Space: O(1)
func canWinNim(n int) bool {
	if n%4 == 0 {
		return false
	}

	return true
}
