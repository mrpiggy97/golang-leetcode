package arrays

func ArrayChange(slice []int) int {
	var previousValue int = -1
	var count int = 0
	for index := range slice {
		var currentVal int = slice[index]
		if index == 0 {
			previousValue = currentVal
		} else {
			if previousValue >= currentVal {
				var newVal int = previousValue + 1
				for secondCount := currentVal; secondCount <= newVal; secondCount++ {
					count += 1
				}
				slice[index] = newVal
			}
		}
	}
	return count
}
