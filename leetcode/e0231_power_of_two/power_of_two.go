package poweroftwo

func isPowerOfTwo(n int) bool {
	if n == 1 {
		return true
	} else if (n <= 0) || ((n % 2) != 0) {
		return false
	} else {
		n = n >> 1
		return isPowerOfTwo(n)
	}
}
