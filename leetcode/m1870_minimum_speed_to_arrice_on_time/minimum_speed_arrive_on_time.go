package minimumspeedtoarriceontime

import "math"

func minSpeedOnTime(dist []int, hour float64) int {
	if len(dist) > int(math.Ceil(hour)) {
		return -1
	}

	maxSpeed := 10_000_000
	minSpeed := 1
	for minSpeed <= maxSpeed {
		speed := minSpeed + (maxSpeed-minSpeed)>>1
		totalHour := 0.0
		for idx, distance := range dist {
			currentHour := float64(distance) / float64(speed)
			if idx == len(dist)-1 {
				totalHour += currentHour
			} else {
				totalHour += math.Ceil(currentHour)
			}
		}

		if totalHour > hour {
			minSpeed = speed + 1
		} else {
			maxSpeed = speed - 1
		}
	}

	return minSpeed
}
