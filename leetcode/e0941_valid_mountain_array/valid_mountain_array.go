package validmountainarray

func validMountainArray(arr []int) bool {
	// two people p1 & p2 are climbing, they will meet each other finally.
	length, p1 := len(arr), 0
	for (p1+1 < length) && (arr[p1] < arr[p1+1]) {
		p1++
	}

	p2 := length - 1
	for (p2-1 >= 0) && (arr[p2] < arr[p2-1]) {
		p2--
	}

	if (p1 > 0) && (p2 < length-1) && (p1 == p2) {
		return true
	}

	return false
}
