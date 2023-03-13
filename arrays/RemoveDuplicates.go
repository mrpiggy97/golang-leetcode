package arrays

func RemoveDuplicates(nums []int) int {
	var count int = 0
	var currentIndex int = 0
	var previousVal int = 0
	for index, value := range nums {
		if index == 0 {
			previousVal = value
			currentIndex = currentIndex + 1
			count += 1
		} else {
			if value != previousVal {
				nums[currentIndex] = value
				previousVal = value
				currentIndex = currentIndex + 1
				count += 1
			}
		}
	}
	return count
}
