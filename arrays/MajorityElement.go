package arrays

func MajorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var counter map[int]int = make(map[int]int)
	var highestCount int = 0
	var selectedValue int = -1
	for _, value := range nums {
		count, exists := counter[value]
		if !exists {
			counter[value] = 1
		} else {
			count += 1
			counter[value] = count
			if count > highestCount {
				highestCount = count
				selectedValue = value
			}
		}
	}
	return selectedValue
}
