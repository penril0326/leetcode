package digit

func SumOfDigit(n int) int {
	sum := 0
	for i := n; i != 0; i /= 10 {
		sum += i % 10
	}

	return sum
}
