package missing_numbers

func missingNumbers(arr []int32, brr []int32) []int32 {
	brrFreq := make(map[int32]int)

	for _, value := range brr {
		brrFreq[value]++
	}

	for _, value := range arr {
		if _, isExist := brrFreq[value]; isExist == true {
			brrFreq[value]--
		}
	}

	result := []int32{}
	for i := 1; i <= 10000; i++ {
		if count, isExits := brrFreq[int32(i)]; (isExits == true) && (count > 0) {
			result = append(result, int32(i))
		}
	}

	return result
}
