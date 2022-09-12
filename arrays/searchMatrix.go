package arrays

func SearchMatrix(matrix [][]int, target int) bool {
	var indexOfLikelySlice int = 0
	for index, slice := range matrix {
		if slice[0] <= target {
			indexOfLikelySlice = index
		} else {
			break
		}
	}
	var selectedSlice []int = matrix[indexOfLikelySlice]
	if selectedSlice[len(selectedSlice)-1] < target {
		return false
	}
	for _, value := range selectedSlice {
		if value > target {
			break
		}
		if value == target {
			return true
		}
	}
	return false
}
