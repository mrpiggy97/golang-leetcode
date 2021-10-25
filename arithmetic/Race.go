package arithmetic

import "math"

func Race(racerSpeed1 int, racerSpeed2 int, distance int) [3]int {
	if racerSpeed1 > racerSpeed2 {
		return [3]int{-1, -1, -1}
	}
	var hours float64 = float64(distance) / float64((racerSpeed2 - racerSpeed1))
	minutes := ((hours * 100) * 60) / 100
	seconds := math.Floor(((math.Floor(hours*10) / 10) * 3600) / 100)
	return [3]int{int(hours), int(minutes), int(seconds)}
}
