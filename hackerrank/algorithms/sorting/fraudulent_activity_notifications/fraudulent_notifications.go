package fraudulent_notification

func activityNotifications(expenditure []int32, d int32) int32 {
	var notifyCount int32
	countList := make([]int32, 201)
	for i := 0; int32(i) < d; i++ {
		value := expenditure[i]
		countList[value]++
	}

	for i := d; int(i) < len(expenditure); i++ {
		new := expenditure[i]
		old := expenditure[i-d]
		median := getMedian(countList, d)

		if float64(new) >= (median * 2.0) {
			notifyCount++
		}

		countList[old]--
		countList[new]++
	}

	return notifyCount
}

func getMedian(countList []int32, d int32) float64 {
	// odd day
	median := 0.0
	var count int32
	if (d % 2) != 0 {
		for i := 0; i < len(countList); i++ {
			count += countList[i]
			if count > (d / 2) {
				median = float64(i)
				break
			}
		}
	} else {
		m1, m2 := 0, 0
		for i := 0; i < len(countList); i++ {
			count += countList[i]
			if (m1 == 0.0) && count >= (d/2) {
				m1 = i
			}

			if (m2 == 0.0) && count >= (d/2+1) {
				m2 = i
				break
			}
		}

		median = float64(m1+m2) / 2.0
	}

	return median
}
