package addbinary

func addBinary(a string, b string) string {
	// 若轉換成int相加可能會overflow，所以此題不需要轉換
	// 直接用字串的處理就好了
	ia, ib := len(a)-1, len(b)-1
	carry := 0
	result := ""

	for (ia >= 0) || (ib >= 0) || (carry > 0) {
		if ia >= 0 {
			carry += int(a[ia] - '0')
			ia--
		}

		if ib >= 0 {
			carry += int(b[ib] - '0')
			ib--
		}

		if carry%2 == 0 {
			result = "0" + result
		} else {
			result = "1" + result
		}

		carry /= 2
	}

	return result
}
