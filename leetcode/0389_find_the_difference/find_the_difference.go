package findthedifference

func findTheDifference(s string, t string) byte {
	var result rune
	for _, r := range s + t {
		result = result ^ r
	}

	return byte(result)
}
