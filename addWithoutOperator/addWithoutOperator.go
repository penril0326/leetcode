package addWithoutOperator

func AddWithoutOperator(a, b int) int {
	if 0 == a {
		return b
	} else if 0 == b {
		return a
	}

	sum := a ^ b
	carry := (a & b) << 1

	return AddWithoutOperator(sum, carry)
}
