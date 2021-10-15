package arithmetic

func DoubleTheAge(fatherYears int, sonYears int) int {
	var originalFatherYears int = fatherYears
	var originalSonYears int = sonYears
	var numberOfYears int = 0
	for sonYears*2 != fatherYears {
		numberOfYears++
		sonYears++
		fatherYears++
		if fatherYears > 120 {
			fatherYears = originalFatherYears
			sonYears = originalSonYears
			numberOfYears = 0
			break
		}
	}

	for fatherYears/2 != sonYears {
		numberOfYears++
		fatherYears--
		sonYears--
	}

	return numberOfYears
}
