package arrays

func RemoveItemFromSlice(index int, slice []int) []int {
	var newValues []int = slice[0:index]
	if index+1 <= len(slice)-1 {
		newValues = append(newValues, slice[index+1:]...)
	}
	return newValues
}

func SortSlice(slice []int) []int {
	var newSlice []int = []int{}
	var newSliceLength int = len(newSlice)
	var oldSliceLength int = len(slice)
	for newSliceLength != oldSliceLength {
		var lowestValue int = slice[0]
		var indexOfLowestValue = 0

		for index, value := range slice {
			if value < lowestValue {
				lowestValue = value
				indexOfLowestValue = index
			}
		}

		newSlice = append(newSlice, lowestValue)
		slice = RemoveItemFromSlice(indexOfLowestValue, slice)
		newSliceLength = len(newSlice)
	}
	return newSlice
}

func Merge(nums1 []int, m int, nums2 []int, n int) {
	for _, value := range nums2 {
		for i, val := range nums1 {
			if val == 0 {
				nums1[i] = value
				break
			}
		}
	}

	var sliceCopy []int = make([]int, len(nums1))
	copy(sliceCopy, nums1)
	sliceCopy = SortSlice(sliceCopy)
	for index := range nums1 {
		nums1[index] = sliceCopy[index]
	}
}
