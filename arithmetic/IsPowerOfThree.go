package arithmetic

import "math"

func IsPowerOfThree(num int) bool {
	for i := 0; i <= 19; i++ {
		var result int = int(math.Pow(3, float64(i)))
		if result == num {
			return true
		}
	}
	return false
}
