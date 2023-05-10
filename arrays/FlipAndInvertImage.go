package arrays

func reverseSlice(slice []int) []int {
	var newSlice []int = []int{}
	for i := len(slice) - 1; i >= 0; i-- {
		newSlice = append(newSlice, slice[i])
	}
	return newSlice
}

func flipSlice(slice []int) []int {
	var newSlice []int = []int{}
	for _, val := range slice {
		if val == 0 {
			newSlice = append(newSlice, 1)
		} else {
			newSlice = append(newSlice, 0)
		}
	}
	return newSlice
}

func FlipAndInvertImage(image [][]int) [][]int {
	var newImage [][]int = [][]int{}
	for _, slice := range image {
		slice = reverseSlice(slice)
		slice = flipSlice(slice)
		newImage = append(newImage, slice)
	}
	return newImage
}
