package array_2d

func hourglassSum(arr [][]int32) int32 {
	if len(arr) < 3 {
		return 0
	}

	for _, subArr := range arr {
		if len(subArr) < 3 {
			return 0
		}
	}

	hourGlassMaxWigth := 3
	hourGlassMaxHeight := 3

	result := -999999
	for i := 0; i < len(arr)-hourGlassMaxWigth+1; i++ {
		for j := 0; j < len(arr[i])-hourGlassMaxHeight+1; j++ {
			sum := arr[i][j] + arr[i][j+1] + arr[i][j+2] + arr[i+1][j+1] + arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
			if sum > int32(result) {
				result = int(sum)
			}
		}
	}

	return int32(result)
}
