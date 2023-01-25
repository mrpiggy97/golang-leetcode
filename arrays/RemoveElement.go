package arrays

func RemoveElement(nums []int, val int) int {
	var currentIndex int = 0
	var count int = 0
	for _, value := range nums {
		if value != val {
			nums[currentIndex] = value
			currentIndex += 1
			count += 1
		}
	}
	return count
}
