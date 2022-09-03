package arrays

func Contains[T comparable](slice []T, value T) (contains bool, count int) {
	contains = false
	count = 0
	for _, val := range slice {
		if val == value {
			if !contains {
				contains = true
			}
			count = count + 1
		}
	}
	return contains, count
}

func CountElement[T comparable](slice []T, value T) (count int) {
	count = 0
	for _, val := range slice {
		if val == value {
			count = count + 1
		}
	}
	return count
}

func GetElementsInCommon[T comparable](slice1 []T, slice2 []T) []T {
	var elements []T = []T{}
	if len(slice1) < len(slice2) {
		for _, value := range slice1 {
			valueInSlice2, _ := Contains[T](slice2, value)
			if valueInSlice2 {
				valueAlreadyInElements, _ := Contains[T](elements, value)
				if !valueAlreadyInElements {
					elements = append(elements, value)
				}
			}
		}
	} else {
		for _, value := range slice2 {
			valueInSlice1, _ := Contains[T](slice1, value)
			if valueInSlice1 {
				valueAlreadyInElements, _ := Contains[T](elements, value)
				if !valueAlreadyInElements {
					elements = append(elements, value)
				}
			}
		}
	}
	return elements
}

func Intersect(nums1 []int, num2 []int) []int {
	var intersection []int = []int{}
	var elementsInCommon []int = GetElementsInCommon[int](nums1, num2)
	for _, value := range elementsInCommon {
		var countInNums1 int = CountElement[int](nums1, value)
		var countInNum2 int = CountElement[int](num2, value)
		var countUsed int = 0

		if countInNums1 != countInNum2 {
			if countInNums1 < countInNum2 {
				countUsed = countInNums1
			} else {
				countUsed = countInNum2
			}
		}
		if countInNums1 == countInNum2 {
			// you can either use countInNums1 or countInNum2
			// if they are of the same value
			countUsed = countInNums1
		}
		for i := 0; i < countUsed; i++ {
			intersection = append(intersection, value)
		}
	}
	return intersection
}
