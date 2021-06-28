package sherlock_array

func balancedSums(arr []int32) string {
	if len(arr) == 1 {
		return "YES"
	}

	leftSum, rightSum := int32(0), sumArr(arr)
	for i := 0; i < len(arr); i++ {
		rightSum -= arr[i]

		if leftSum < rightSum {
			leftSum += arr[i]
		} else if leftSum == rightSum {
			return "YES"
		}
	}

	return "NO"
}

func sumArr(arr []int32) int32 {
	var sum int32
	for _, v := range arr {
		sum += v
	}

	return sum
}
